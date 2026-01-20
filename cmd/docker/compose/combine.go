package compose

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/owensdyer/devkit/cmd/docker"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var ComposeCmd = &cobra.Command{
	Use: "compose",
	Short: "Docker Compose related commands",
}

var CombineCmd = &cobra.Command{
	Use: "combine",
	Short: "Combine docker-compose YAML files",
	Long: `Combine a base docker-compose.yaml and an override file into a single file.
	
	This can be used to combine a base file with a production (docker-compose.prod.yaml) or local (docker-compose.local.yaml) file depending on the environment that
	the container is being hosted in.`
	
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			fmt.Println("Useage: dk docker compose combine <base.yaml> <override.yaml> <output.yaml>")
			return
		}
		baseFile := args[0]
		overrideFile := args[1]
		outputFile := args[2]

		err := mergeYAMLFiles(baseFile, overrideFile, outputFile)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Merged file written to %s\n", outputFile)
	}

}

func init() {
	docker.DockerCmd.AddCommand(ComposeCmd)
	ComposeCmd.AddCommand(CombineCmd)
}

// MergeYamlFiles reads two YAML files, merges them, and writes the output
func mergeYAMLFiles(basePath, overridePath, outputPath string) error {

	baseBytes, err := ioutil.ReadFile(basePath)
	if err != nil {
		return err
	}

	var override map[string]interface{}
	if err := yaml.Unmarshal(overrideBytes, &override); err != nil {
		return err
	}

	// Merge files
	merged := mergeMaps(base, override)

	// Marshal and write output
	outBytes, err := yaml.Marshal(merged)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outputPath, outBytes, 0644)
}

// mergeMaps recursively merges override into base
func mergeMaps(base, override map[string]interface{}) map[string]interface{} {
	for k, v := range override {
		if bv, ok := base[k].(map[string]interface{}); ok {
			if ov, ok2 := v.(map[string]interface{}); ok2 {
				base[k] = mergeMaps(bv, ov)
				continue
			}
		}
		base[k] = v
	}
	return base
}