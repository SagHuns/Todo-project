package configs

import "github.com/spf13/viper"

// A ideia é fazer com que ninguém, fora o package configs, consiga acessar as configurações diretamente
// e também sem poder alterar o valor, o package apenas dá o valor.
var cfg *config

// Une os dois configs em uma struct só para facilitar o acesso
type config struct {
	API APIconfig
	DB DBconfig
}

type APIconfig struct {
	Port string
}

type DBconfig struct {
	Host string
	Port string
	User string
	Pass string
	Database string
}

// Sempre chamada no start das aplicações, definindo alguns valores padrões
func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
}

func Load() error {
	viper.SetConfigName("config")  // Nome do arquivo que ele procurará
	viper.SetConfigType("toml")  // Extensão do arquivo procurado
	viper.AddConfigPath(".")  // Path do arquivo
	err := viper.ReadInConfig()  // Mandando o viper ler o arquivo que setamos as configurações acima
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)  // Alocando espaço na memória (como se utilizasse ponteiro)

	cfg.API = APIconfig{
		Port: viper.GetString("api.port"),
	}
	
	cfg.DB = DBconfig{
		Host: viper.GetString("database.host"),
		Port: viper.GetString("database.port"),
		User: viper.GetString("database.user"),
		Pass: viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}
	return nil
}

// Getters

func GetDB() DBconfig {
	return cfg.DB
} 

func GetServerPort() string {
	return cfg.API.Port
}