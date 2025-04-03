# Gorder Dir:
app：一般放应用层相关代码和启动文件，也可能包含部分基础设施初始化。

adapters: 适配器，用来适配不同的底层存储,用来进行CRUD操作,类似MVC 中的models part,但其是依赖倒置（依赖于接口而非具体实现）而非依赖ORM

domain：核心领域模型、实体、领域服务等，体现业务规则和概念。

ports：系统对外或对内的端口接口（inbound/outbound），定义与外部系统或接口适配器交互的契约。

service：应用服务层，用来编排领域对象或实现用例逻辑。

tmp：临时文件或开发调试用目录。

order / payment / stock / kitchen：代表不同的业务子域或上下文。每个目录里可能有各自的领域、应用逻辑。

scripts：脚本文件，管理构建、部署、迁移等流程。