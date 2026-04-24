package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	goyaml "github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
	"github.com/unbound-force/unbound-force/internal/config"
)

// --- Init subcommand ---

type configInitParams struct {
	targetDir string
	stdout    io.Writer
}

func runConfigInit(p configInitParams) error {
	result, err := config.InitFile(config.InitOptions{
		ProjectDir: p.targetDir,
	})
	if err != nil {
		return err
	}

	if result.Created {
		fmt.Fprintf(p.stdout, "Created %s\n", result.Path)
		fmt.Fprintln(p.stdout, "All values are commented out — uncomment what you want to change.")
		return nil
	}

	if result.Updated {
		fmt.Fprintf(p.stdout, "Updated %s\n", result.Path)
		if len(result.SectionsAdded) > 0 {
			fmt.Fprintf(p.stdout, "  Added sections: %v\n", result.SectionsAdded)
		}
		if len(result.SectionsRemoved) > 0 {
			fmt.Fprintf(p.stdout, "  Removed deprecated sections: %v\n", result.SectionsRemoved)
		}
		fmt.Fprintln(p.stdout, "A backup was saved to .uf/config.yaml.bak")
		return nil
	}

	fmt.Fprintf(p.stdout, "%s is already up to date.\n", result.Path)
	return nil
}

// --- Show subcommand ---

type configShowParams struct {
	targetDir string
	format    string
	stdout    io.Writer
}

func runConfigShow(p configShowParams) error {
	cfg, err := config.Load(config.LoadOptions{
		ProjectDir: p.targetDir,
	})
	if err != nil {
		return err
	}

	switch p.format {
	case "json":
		data, jsonErr := json.MarshalIndent(cfg, "", "  ")
		if jsonErr != nil {
			return fmt.Errorf("marshal JSON: %w", jsonErr)
		}
		fmt.Fprintln(p.stdout, string(data))
	default:
		data, yamlErr := goyaml.Marshal(cfg)
		if yamlErr != nil {
			return fmt.Errorf("marshal YAML: %w", yamlErr)
		}
		fmt.Fprint(p.stdout, string(data))
	}
	return nil
}

// --- Validate subcommand ---

type configValidateParams struct {
	targetDir string
	format    string
	stdout    io.Writer
}

func runConfigValidate(p configValidateParams) error {
	configPath := config.RepoConfigPath(p.targetDir)
	data, err := os.ReadFile(configPath)
	if err != nil {
		// Missing file is valid — defaults are used.
		fmt.Fprintln(p.stdout, "No config file found at", configPath)
		fmt.Fprintln(p.stdout, "This is valid — compiled defaults are used.")
		return nil
	}

	// Try to parse the YAML.
	var cfg config.Config
	if parseErr := goyaml.Unmarshal(data, &cfg); parseErr != nil {
		fmt.Fprintf(p.stdout, "FAIL: %s is not valid YAML\n", configPath)
		fmt.Fprintf(p.stdout, "  Error: %v\n", parseErr)
		return fmt.Errorf("config validation failed")
	}

	// Validate known field values.
	var errors []string
	errors = append(errors, validateSetup(cfg.Setup)...)
	errors = append(errors, validateSandbox(cfg.Sandbox)...)
	errors = append(errors, validateGateway(cfg.Gateway)...)
	errors = append(errors, validateEmbedding(cfg.Embedding)...)
	errors = append(errors, validateDoctor(cfg.Doctor)...)

	if len(errors) > 0 {
		fmt.Fprintf(p.stdout, "Config validation: %s\n\n", configPath)
		for _, e := range errors {
			fmt.Fprintf(p.stdout, "  FAIL: %s\n", e)
		}
		fmt.Fprintf(p.stdout, "\n%d error(s) found\n", len(errors))
		return fmt.Errorf("config validation failed")
	}

	fmt.Fprintf(p.stdout, "Config validation: %s\n", configPath)
	fmt.Fprintln(p.stdout, "  All checks passed.")
	return nil
}

// --- Field validators ---

func validateSetup(cfg config.SetupConfig) []string {
	var errs []string
	valid := map[string]bool{
		"auto": true, "homebrew": true, "dnf": true,
		"apt": true, "manual": true, "": true,
	}
	if !valid[cfg.PackageManager] {
		errs = append(errs, fmt.Sprintf(
			"setup.package_manager: %q is not valid (auto|homebrew|dnf|apt|manual)",
			cfg.PackageManager))
	}

	validMethods := map[string]bool{
		"auto": true, "homebrew": true, "dnf": true, "rpm": true,
		"apt": true, "curl": true, "skip": true, "nvm": true,
		"fnm": true, "mise": true, "": true,
	}
	for name, tool := range cfg.Tools {
		if !validMethods[tool.Method] {
			errs = append(errs, fmt.Sprintf(
				"setup.tools.%s.method: %q is not valid",
				name, tool.Method))
		}
	}
	return errs
}

func validateSandbox(cfg config.SandboxConfig) []string {
	var errs []string
	validRuntime := map[string]bool{
		"auto": true, "podman": true, "docker": true, "": true,
	}
	if !validRuntime[cfg.Runtime] {
		errs = append(errs, fmt.Sprintf(
			"sandbox.runtime: %q is not valid (auto|podman|docker)",
			cfg.Runtime))
	}

	validBackend := map[string]bool{
		"auto": true, "podman": true, "che": true, "": true,
	}
	if !validBackend[cfg.Backend] {
		errs = append(errs, fmt.Sprintf(
			"sandbox.backend: %q is not valid (auto|podman|che)",
			cfg.Backend))
	}

	validMode := map[string]bool{
		"isolated": true, "direct": true, "": true,
	}
	if !validMode[cfg.Mode] {
		errs = append(errs, fmt.Sprintf(
			"sandbox.mode: %q is not valid (isolated|direct)",
			cfg.Mode))
	}
	return errs
}

func validateGateway(cfg config.GatewayConfig) []string {
	var errs []string
	validProvider := map[string]bool{
		"auto": true, "anthropic": true, "vertex": true,
		"bedrock": true, "": true,
	}
	if !validProvider[cfg.Provider] {
		errs = append(errs, fmt.Sprintf(
			"gateway.provider: %q is not valid (auto|anthropic|vertex|bedrock)",
			cfg.Provider))
	}
	if cfg.Port < 0 || cfg.Port > 65535 {
		errs = append(errs, fmt.Sprintf(
			"gateway.port: %d is not a valid port number (0-65535)",
			cfg.Port))
	}
	return errs
}

func validateEmbedding(cfg config.EmbeddingConfig) []string {
	var errs []string
	validProvider := map[string]bool{
		"ollama": true, "": true,
	}
	if !validProvider[cfg.Provider] {
		errs = append(errs, fmt.Sprintf(
			"embedding.provider: %q is not valid (ollama)",
			cfg.Provider))
	}
	if cfg.Dimensions < 0 {
		errs = append(errs, fmt.Sprintf(
			"embedding.dimensions: %d must be non-negative",
			cfg.Dimensions))
	}
	return errs
}

func validateDoctor(cfg config.DoctorConfig) []string {
	var errs []string
	validSeverity := map[string]bool{
		"required": true, "recommended": true, "optional": true,
	}
	for name, sev := range cfg.Tools {
		if !validSeverity[sev] {
			errs = append(errs, fmt.Sprintf(
				"doctor.tools.%s: %q is not valid (required|recommended|optional)",
				name, sev))
		}
	}
	return errs
}

// --- Command factory ---

func newConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage Unbound Force configuration",
		Long: `Manage the unified .uf/config.yaml configuration file.

Subcommands:
  init      Create or update the config file
  show      Display effective config after all layers merge
  validate  Validate config against known field values`,
	}

	// --- init subcommand ---
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Create or update .uf/config.yaml",
		RunE: func(cmd *cobra.Command, _ []string) error {
			dir, _ := cmd.Flags().GetString("dir")
			return runConfigInit(configInitParams{
				targetDir: dir,
				stdout:    cmd.OutOrStdout(),
			})
		},
	}
	initCmd.Flags().String("dir", ".", "Target directory")

	// --- show subcommand ---
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "Display effective configuration",
		RunE: func(cmd *cobra.Command, _ []string) error {
			dir, _ := cmd.Flags().GetString("dir")
			format, _ := cmd.Flags().GetString("format")
			return runConfigShow(configShowParams{
				targetDir: dir,
				format:    format,
				stdout:    cmd.OutOrStdout(),
			})
		},
	}
	showCmd.Flags().String("dir", ".", "Target directory")
	showCmd.Flags().String("format", "text", "Output format (text|json)")

	// --- validate subcommand ---
	validateCmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate config file",
		RunE: func(cmd *cobra.Command, _ []string) error {
			dir, _ := cmd.Flags().GetString("dir")
			format, _ := cmd.Flags().GetString("format")
			return runConfigValidate(configValidateParams{
				targetDir: dir,
				format:    format,
				stdout:    cmd.OutOrStdout(),
			})
		},
	}
	validateCmd.Flags().String("dir", ".", "Target directory")
	validateCmd.Flags().String("format", "text", "Output format (text|json)")

	cmd.AddCommand(initCmd)
	cmd.AddCommand(showCmd)
	cmd.AddCommand(validateCmd)

	return cmd
}
