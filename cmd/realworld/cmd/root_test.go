package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func TestExecute(t *testing.T) {
	c := qt.New(t)
	got := Execute()
	c.Assert(got, qt.IsNil)
}

func TestConfig(t *testing.T) {
	c := qt.New(t)

	const keyPort = "test.port"

	testCmd := &cobra.Command{
		Use: "test",
		Run: func(_ *cobra.Command, _ []string) {},
	}
	testCmd.Flags().Int("port", defaultServePort, "port")
	_ = viper.BindPFlag(keyPort, testCmd.Flags().Lookup("port"))

	rootCmd.AddCommand(testCmd)
	defer rootCmd.RemoveCommand(testCmd)

	c.Run("default", func(c *qt.C) {
		rootCmd.SetArgs([]string{"test"})
		c.Assert(rootCmd.Execute(), qt.IsNil)

		got := viper.GetInt(keyPort)
		c.Assert(got, qt.Equals, 8080)
	})

	c.Run("config", func(c *qt.C) {
		dir := c.TempDir()
		cfgStr := `
test:
  port: 8081
`
		filename := filepath.Join(dir, "config.yml")
		writeFile(c, filename, []byte(cfgStr))

		rootCmd.SetArgs([]string{"test", "--config", filename})
		c.Assert(rootCmd.Execute(), qt.IsNil)

		got := viper.GetInt(keyPort)
		c.Assert(got, qt.Equals, 8081)

		// reset config
		config = ""
	})

	c.Run("environment", func(c *qt.C) {
		c.Setenv("REALWORLD_TEST_PORT", "8082")

		rootCmd.SetArgs([]string{"test"})
		c.Assert(rootCmd.Execute(), qt.IsNil)

		got := viper.GetInt(keyPort)
		c.Assert(got, qt.Equals, 8082)
	})

	c.Run("flag", func(c *qt.C) {
		c.Setenv("REALWORLD_TEST_PORT", "8082")

		rootCmd.SetArgs([]string{"test", "--port", "8083"})
		c.Assert(rootCmd.Execute(), qt.IsNil)

		got := viper.GetInt(keyPort)
		c.Assert(got, qt.Equals, 8083)
	})
}

func writeFile(c *qt.C, filename string, data []byte) {
	c.Assert(ioutil.WriteFile(filename, data, os.FileMode(0755)), qt.IsNil)
}
