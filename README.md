# dam.example_app

## Для чего DAM ?

* облегчает разработку и тестирование микро сервисных приложений,
  для локального развертывания без привязки к среде разработки,
  а также архитектуре операционной системы

* для развертывания и поддержке продуктов на машинах с ограниченным доступом к интернет

* администраторам и сопровождению, чтобы облегчить дальнейшую настройку и контролировать обновление продукта

## dam help
```
Docker Application Manager

Version:
1.2.1

Usage:
dam [command]

Available Commands:
addrepo     Add an app registry to the system.
completion  generate the autocompletion script for the specified shell
create      Create docker application.
export      Export apps to file.
help        Help about any command
import      Import apps from file.
info        Information for your application.
install     Install docker application from a docker registry or a specific file archive.
list        List all installed your applications.
listrepos   List all defined repositories.
modifyrepo  Modify properties of repositories specified.
purge       Remove docker images not attached to apps.
remove      Remove docker application.
removerepo  Remove registry specified by name or number.
save        Save app to an archive.
search      Search app in remote registry.

Flags:
-x, --debug   Enable debug mode
-h, --help    help for dam

Use "dam [command] --help" for more information about a command.
```
