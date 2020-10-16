package command

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/pflag"
	"github.com/svang/api/internal/config"
	"github.com/svang/api/internal/yaml"
)

var logDirPathFromEnv string //This will be set through the build command, see Makefile

//constants needed
const (
	SampleConfigFileName = "svang-sample.yaml"
	ConfigFileName       = "svang.yaml"
)

// Removes slash from the end
func parsePath(path string) string {
	lastChar := path[len(path)-1:]

	if lastChar != "/" {
		path += "/"
	}
	return path
}

func init() {

	if logDirPathFromEnv != "" {
		logDirPathFromEnv = parsePath(logDirPathFromEnv)
	} else {
		logDirPathFromEnv = "logs/"
	}

	if _, err := os.Stat(logDirPathFromEnv); os.IsNotExist(err) {
		fmt.Printf("\n%v directory does not exist, creating ...\n\n", logDirPathFromEnv)
		err := os.Mkdir(logDirPathFromEnv, 0755)
		if err != nil {
			log.Fatalf("init error while creating %v directory, cannot proceed. err: %v", logDirPathFromEnv, err)
		}
	}
}

//New creates a new instance for command execution
func New(flags *pflag.FlagSet) *Instance {
	return &Instance{
		config:    &config.SvangType{},
		Flags:     flags,
		StartTime: time.Now(),
	}
}

//CreateSampleConfigFile creates sample config file
func (i *Instance) CreateSampleConfigFile() *Instance {
	if i.Error != nil {
		return i
	}
	err := yaml.CreateSampleConfigFile(SampleConfigFileName)
	if err != nil {
		i.Error = err
	}
	fmt.Println("\nGenerated sample file : ", SampleConfigFileName)
	fmt.Println("Copy", SampleConfigFileName, "to", ConfigFileName, "and make edits as per your environment.")
	return i
}

//TimeTaken prints time taken for execution
func (i *Instance) TimeTaken() *Instance {
	if i.Error != nil {
		return i
	}
	fmt.Println("Time taken : ", time.Since(i.StartTime))
	return i
}

// Log file to display part size and elapsed time
func Log(file string, msg string) {
	var logpath = file
	f, err := os.OpenFile(logpath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
		return
	}

	defer f.Close()

	//set output of logs to f
	log.SetOutput(f)

	log.Println(msg)
}
