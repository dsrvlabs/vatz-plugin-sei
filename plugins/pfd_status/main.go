package main

import (
	"flag"
	"fmt"
	"os/exec"

	pluginpb "github.com/dsrvlabs/vatz-proto/plugin/v1"
	"github.com/dsrvlabs/vatz/sdk"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/structpb"
)

const (

	// Default values.
	defaultRPCAddr = "http://localhost:1317"
	defaultAddr    = "127.0.0.1"
	defaultPort    = 10004
	pluginName     = "pfd_status"
)

var (
	rpcAddr     string
	addr        string
	port        int
	valoperAddr string
)

func init() {
	flag.StringVar(&rpcAddr, "rpcURI", defaultRPCAddr, "CosmosHub RPC URI Address")
	flag.StringVar(&addr, "addr", defaultAddr, "Listening address")
	flag.IntVar(&port, "port", defaultPort, "Listening port")
	flag.StringVar(&valoperAddr, "valoperAddr", "", "CosmosHub validator operator address")

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
	severity := pluginpb.SEVERITY_INFO
	state := pluginpb.STATE_NONE

	var msg string

	if err := checkSeidInstallation(); err != nil {
		severity = pluginpb.SEVERITY_CRITICAL
		state = pluginpb.STATE_FAILURE
		msg = "Failed to get price-feeder status"
		log.Info().Str("moudle", "plugin").Msg(msg)

		return sdk.CallResponse{}, err
	}

	// Set up the command and arguments
	cmd := exec.Command("seid", "q", "oracle", "vote-penalty-counter", valoperAddr)

	// Execute the command and collect the result
	output, err := cmd.CombinedOutput()
	if err != nil {
		return sdk.CallResponse{}, fmt.Errorf("Error occurred while executing the command: %v\nError Details: %s", err, output)
	}

	severity = pluginpb.SEVERITY_INFO
	msg = fmt.Sprintf("Validator bonded. included active set")
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
