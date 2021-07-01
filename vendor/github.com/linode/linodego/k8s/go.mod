module github.com/azurerm/azurermgo/k8s

require (
	github.com/azurerm/azurermgo v0.20.1
	k8s.io/api v0.18.3
	k8s.io/apimachinery v0.18.3
	k8s.io/client-go v0.18.3
)

replace github.com/azurerm/azurermgo => ../

go 1.13
