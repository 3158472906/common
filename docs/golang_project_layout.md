<h1>
<a href="https://studygolang.com/articles/26941?fr=sidebar">Article source</a>
</h1>

#Go 项目标准布局（结构）
>预计阅读时间 2 分钟 · 不到1分钟
>这是一个创建于 2020-03-03 15:33:24 的文章，其中的信息可能已经有所发展或是发生改变。
> ```https://github.com/golang-standards/project-layout``` 项目总结了 Go 项目的布局，这是Go应用程序项目的基本布局，并不是Go核心开发团队定义的官方标准；但是，这是Go生态系统中的一种常见的项目布局模式。

#Go目录
##/cmd
>该目录用于存放 Go 项目的入口，即 main.go。一般来说，我们应该在 cmd 目录下创建子目录，子目录名称代表可执行程序的名称（例如/cmd/myapp）。上面列出的优秀开源项目基本上遵循了这一规则。
>一般来说，该目录中的代码应该尽可能少。如果认为该代码可以导入并在其他项目中使用，那么它应该位于/pkg目录中。如果该代码不可重用，或者不希望其他人重用它，则将该代码放在/internal目录中。

##/internal
>这是 Go 包的一个特性，放在该包中的代码，表明只希望项目内部使用，是项目或库私有的，其他项目或库不能使用。请注意，不限于顶层internal目录，internal在项目树的任何级别上都可以有多个目录。
>可以选择向内部包中添加一些额外的结构，以分隔共享和非共享内部代码。它不是必需的（尤其是对于较小的项目），但是最好有视觉提示来显示包的用途。实际应用程序代码可以进入/internal/app目录（例如/internal/app/myapp），而这些应用程序共享的代码可以进入/internal/pkg目录（例如/internal/pkg/myprivlib）。

##/pkg
>该包可以和 internal 对应，是公开的。一般来说，放在该包的代码应该和具体业务无关，方便本项目和其他项目重用。当你决定将代码放入该包时，你应该对其负责，因为别人很可能使用它。
>如果应用程序项目很小，并且嵌套的额外层次不会增加太多价值（除非您真的想要，请不要使用它。当它变得足够大并且您的根目录变得非常复杂时（特别是如果您有很多非Go应用程序组件），请考虑一下。

##/vendor
>应用程序依存关系，GO1.13启用go module来代替，并不需要vendor目录了。

>server application目录
##api
>该目录用来存放 OpenAPI/Swagger 规则说明, JSON 格式定义, 协议定义文件等。也有可能用来存放具体的对外公开 API.

>web application 目录
##/web
>Web应用程序特定的组件：静态Web资产，服务器端模板和SPA。

>common application目录
##/configs
>配置文件模板或默认配置。

##/init
>存放随着系统自动启动脚本，如：systemd, upstart, sysv；或者通过 supervisor 进行进程管理的脚本。

##/scripts
>存放 build、install、analysis 等操作脚本。这些脚本使得项目根目录的 Makefile 很简洁。

##/build
>该目录用于存放打包和持续集成相关脚本。将云（AMI），容器（Docker），操作系统（deb，rpm，pkg）软件包配置和脚本放在/build/package目录中。
>将CI（travis，circle，drone）配置和脚本放在/build/ci目录中。请注意，某些配置项工具（例如Travis CI）对于其配置文件的位置非常挑剔。尝试将配置文件放在/build/ci目录中，将它们链接到CI工具期望它们的位置（如果可能）。

##/deployments
>IaaS，PaaS，系统和容器编排部署配置和模板（docker-compose，kubernetes /
> helm，mesos，terraform，bosh）。

##/test
>一般用来存放除单元测试、基准测试之外的测试，比如集成测试、测试数据等。

>其他目录
##/docs
>设计和用户文档（除了godoc生成的文档之外）。

##/tools
>存放项目的支持工具。请注意，这些工具可以从/pkg和/internal目录导入代码。

##/examples
>应用程序或公共库的示例。

##/third_party
>外部帮助程序工具，分叉的代码和其他第三方工具（例如Swagger UI）。

##/githooks
>githooks

##/assets
>与资源库一起使用的其他资产（图像，徽标等）。

##/website
>如果不使用Github页面，则在这里放置项目的网站数据。

>不应该拥有的目录
>/src