package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	pluginpb "github.com/dsrvlabs/vatz-proto/plugin/v1"
	"github.com/dsrvlabs/vatz/sdk"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	// Default values.
	defaultRPCAddr  = "http://localhost:1317"
	defaultAddr     = "127.0.0.1"
	defaultPort     = 10004
	defaultWarn     = 20
	defaultCritical = 95
	pluginName      = "pfd_status"
)

var (
	rpcAddr           string
	addr              string
	port              int
	valoperAddr       string
	warnCondition     float64
	criticalCondition float64
	seiHome           string
)

func init() {
	flag.StringVar(&rpcAddr, "rpcURI", defaultRPCAddr, "CosmosHub RPC URI Address")
	flag.StringVar(&addr, "addr", defaultAddr, "Listening address")
	flag.IntVar(&port, "port", defaultPort, "Listening port")
	flag.StringVar(&valoperAddr, "valoperAddr", "", "CosmosHub validator operator address")
	flag.Float64Var(&warnCondition, "port", defaultWarn, "Warning count")
	flag.Float64Var(&criticalCondition, "port", defaultCritical, "Critical count")
	flag.StringVar(&seiHome, "seiHome", "", "Sei node's home flag")

	flag.Parse()
}

func main() {
	if valoperAddr == "" {
		log.Fatal().Str("module", "plugin").Msg("Please specify -valoperAddr")
	}

	p := sdk.NewPlugin(pluginName)
	p.Register(pluginFeature)

	ctx := context.Background()
	if err := p.Start(ctx, addr, port); err != nil {
		fmt.Println("exit")
	}
}

func checkSeidInstallation() error {
	// Find the path to the seid command
	_, err := exec.LookPath("seid")
	if err != nil {
		return fmt.Errorf("seid command not found. Please install seid.")
	}

	return nil
}

func pluginFeature(info, option map[string]*structpb.Value) (sdk.CallResponse, error) {
	severity := pluginpb.SEVERITY_ERROR
	state := pluginpb.STATE_FAILURE

	var msg string

	if err := checkSeidInstallation(); err != nil {
		msg = "Failed to get price-feeder status"
		log.Info().Str("moudle", "plugin").Msg(msg)

		return sdk.CallResponse{}, err
	}

	// Set up the command and arguments
	cmd := exec.Command("seid", "q", "oracle", "vote-penalty-counter", valoperAddr, "--home", seiHome)

	output, err := cmd.CombinedOutput()
	if err != nil {
		msg = fmt.Sprintf("Error occurred:", err)
	}

	// Print the original result.
	//	fmt.Println("Original Result:")
	//	fmt.Println(string(output))

	lines := strings.Split(string(output), "\n")
	votePenaltyCounter := make(map[string]string)

	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			votePenaltyCounter[key] = strings.Trim(value, "\"")
		}
	}

	// Extract the necessary information.
	abstainCountStr := votePenaltyCounter["abstain_count"]
	successCountStr := votePenaltyCounter["success_count"]

	abstainCount, err := strconv.Atoi(abstainCountStr)
	if err != nil {
		msg = fmt.Sprintf("Error parsing abstain_count:", err)
	}

	successCount, err := strconv.Atoi(successCountStr)
	if err != nil {
		msg = fmt.Sprintf("Error parsing success_count:", err)
	}

	missingRatio := float64(abstainCount) / float64(successCount) * 100

	if missingRatio > criticalCondition {
		severity = pluginpb.SEVERITY_CRITICAL
		state = pluginpb.STATE_SUCCESS
		msg = fmt.Sprintf("Price-Feeder missing rate is too high: %.2f%%\nYou're going to gt jailed.\n", missingRatio)
	} else if missingRatio > warnCondition {
		severity = pluginpb.SEVERITY_WARNING
		state = pluginpb.STATE_SUCCESS
		msg = fmt.Sprintf("Price-Feeder oracle missing rate are rising: %.2f%%\n", missingRatio)
	} else {
		severity = pluginpb.SEVERITY_INFO
		state = pluginpb.STATE_SUCCESS
		msg = fmt.Sprintf("Price-Feeder oracle missing rate is good: %.2f%%\n", missingRatio)
	}

	log.Debug().Str("module", "plugin").Msg(msg)

	ret := sdk.CallResponse{
		FuncName:   info["execute_method"].GetStringValue(),
		Message:    msg,
		Severity:   severity,
		State:      state,
		AlertTypes: []pluginpb.ALERT_TYPE{pluginpb.ALERT_TYPE_DISCORD},
	}

	return ret, nil

}
