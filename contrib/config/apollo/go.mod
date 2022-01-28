module github.com/go-kratos/kratos/contrib/config/apollo/v2

go 1.16

require (
	github.com/apolloconfig/agollo/v4 v4.0.8
	github.com/go-kratos/kratos/v2 v2.1.5
)

require (
	github.com/spf13/afero v1.8.0 // indirect
	github.com/spf13/viper v1.10.1 // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	gopkg.in/ini.v1 v1.66.3 // indirect
)

replace github.com/go-kratos/kratos/v2 => ../../../
