package cli

import "github.com/spf13/cobra"

type CLI struct {
	RootCommand *cobra.Command
}

type StringFlag struct {
	Name         string
	ShortName    string
	DefaultValue string
	Usage        string
}

type IntFlag struct {
	Name         string
	ShortName    string
	DefaultValue int
	Usage        string
}

type BoolFlag struct {
	Name         string
	ShortName    string
	DefaultValue bool
	Usage        string
}

func (cli *CLI) AddCommand(commandName string, shortDescription string, runFunc func(*CLI)) {
	cli.RootCommand = &cobra.Command{
		Use:   commandName,
		Short: shortDescription,
		Run: func(cmd *cobra.Command, args []string) {
			runFunc(cli)
		},
	}
}

func (cli *CLI) AddBoolFlag(b *BoolFlag) {
	cli.RootCommand.PersistentFlags().BoolP(b.Name, b.ShortName, b.DefaultValue, b.Usage)
}

func (cli *CLI) AddIntFlag(i *IntFlag) {
	cli.RootCommand.PersistentFlags().IntP(i.Name, i.ShortName, i.DefaultValue, i.Usage)
}

func (cli *CLI) AddStringFlag(i *StringFlag) {
	cli.RootCommand.PersistentFlags().StringP(i.Name, i.ShortName, i.DefaultValue, i.Usage)
}

func (cli *CLI) Run() {
	cli.RootCommand.Execute()
}

func (cli *CLI) GetFlagValue(flagName string) string {
	return cli.RootCommand.Flag(flagName).Value.String()
}
