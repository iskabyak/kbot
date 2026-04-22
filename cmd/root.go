var rootCmd = &cobra.Command{
	Use:   "kbot",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application.`,

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Файл README.md існує.")

		data, err := os.ReadFile("README.md")
		if err != nil {
			return
		}

		content := string(data)

		if strings.Contains(content, "TELE_TOKEN") {
			fmt.Println("This project uses TELE_TOKEN environment variable.")
		} else if strings.Contains(content, ".Context") {
			fmt.Println(".Context exists")
		}
	},
}
