在DDD架构中，防腐层（Anti-Corruption Layer, ACL）是一种用于解耦上下文之间模型的边界模式。它能够将不同子系统或限界上下文的通信协议、数据模型和领域语言相互转换，从而保护核心领域模型免受外部系统的污染。防腐层最初由Eric Evans在《Domain-Driven Design》一书中提出，常通过外观（Facade）、适配器（Adapter）和数据传输对象（DTO）等设计模式来实现。

## 防腐层是什么
### 定义
**防腐层** 是一种架构模式，用于在两个模型或系统之间建立中介层，翻译它们的通信，使每个子系统都能使用自己的领域语言和模型进行交互([learn.microsoft.com](https://learn.microsoft.com/en-us/azure/architecture/patterns/anti-corruption-layer?utm_source=chatgpt.com), [codeopinion.com](https://codeopinion.com/anti-corruption-layer-for-mapping-between-boundaries/?utm_source=chatgpt.com))。该模式最早由Eric Evans在DDD中提出([learn.microsoft.com](https://learn.microsoft.com/en-us/azure/architecture/patterns/anti-corruption-layer?utm_source=chatgpt.com))。

### 核心作用
- **隔离依赖**：确保新系统的设计不受遗留系统或外部系统的限制和质量问题影响([learn.microsoft.com](https://learn.microsoft.com/en-us/azure/architecture/patterns/anti-corruption-layer?utm_source=chatgpt.com), [cloudbees.com](https://www.cloudbees.com/blog/anti-corruption-layer-how-keep-legacy-support-breaking-new-systems?utm_source=chatgpt.com))。
- **模型保护**：防止外部上下文的领域模型污染或破坏核心领域模型的完整性和一致性([de.wikipedia.org](https://de.wikipedia.org/wiki/Domain-driven_Design?utm_source=chatgpt.com), [codeopinion.com](https://codeopinion.com/anti-corruption-layer-for-mapping-between-boundaries/?utm_source=chatgpt.com))。
- **语言适配**：将一个上下文的泛化语言（Ubiquitous Language）映射到另一个上下文的专用术语，实现无缝集成([cloudbees.com](https://www.cloudbees.com/blog/anti-corruption-layer-how-keep-legacy-support-breaking-new-systems?utm_source=chatgpt.com), [de.wikipedia.org](https://de.wikipedia.org/wiki/Domain-driven_Design?utm_source=chatgpt.com))。

## 场景与使用时机
### 适用场景
1. **多个限界上下文通信**：当不同上下文拥有各自模型和语言，需要相互调用时([ddd-practitioners.com](https://ddd-practitioners.com/home/glossary/bounded-context/bounded-context-relationship/anticorruption-layer/?utm_source=chatgpt.com))。
2. **与遗留系统集成**：新系统需要调用功能落后、数据模式复杂的旧系统时([learn.microsoft.com](https://learn.microsoft.com/en-us/azure/architecture/patterns/anti-corruption-layer?utm_source=chatgpt.com), [docs.aws.amazon.com](https://docs.aws.amazon.com/prescriptive-guidance/latest/cloud-design-patterns/acl.html?utm_source=chatgpt.com))。
3. **外部服务依赖**：与第三方API或外部微服务交互时，防止外部契约直接入侵本地模型([ddd-practitioners.com](https://ddd-practitioners.com/home/glossary/bounded-context/bounded-context-relationship/anticorruption-layer/?utm_source=chatgpt.com))。

### 不适用场景
- 当上下文间数据模型高度一致，额外的中介层带来过度复杂度时，应评估是否必要([learn.microsoft.com](https://learn.microsoft.com/en-us/azure/architecture/patterns/anti-corruption-layer?utm_source=chatgpt.com), [ddd-practitioners.com](https://ddd-practitioners.com/home/glossary/bounded-context/bounded-context-relationship/anticorruption-layer/?utm_source=chatgpt.com))。

## 实现与模式
### 常用设计模式
- **外观（Facade）**：为外部系统提供统一简化接口。
- **适配器（Adapter）**：将外部接口转换为本地接口。
- **数据传输对象（DTO）**：表示跨上下文传输的载体对象([de.wikipedia.org](https://de.wikipedia.org/wiki/Domain-driven_Design?utm_source=chatgpt.com))。

### 部署方式
- **内嵌式**：防腐层作为应用内部组件实现。
- **独立服务**：将防腐层独立部署为微服务，便于伸缩和管理([docs.aws.amazon.com](https://docs.aws.amazon.com/prescriptive-guidance/latest/cloud-design-patterns/acl.html?utm_source=chatgpt.com), [handbook.chaineapp.com](https://handbook.chaineapp.com/handbook/engineering/resources/ddd/domain-driven-design-handbook/?utm_source=chatgpt.com))。

### 设计注意事项
- **性能影响**：中介层可能增加调用延迟。
- **运维耗费**：需额外维护监控、部署和一致性保证([learn.microsoft.com](https://learn.microsoft.com/en-us/azure/architecture/patterns/anti-corruption-layer?utm_source=chatgpt.com))。
- **演进策略**：可将防腐层视为迁移方案的一部分，待旧系统完全迁移后再逐步剔除([learn.microsoft.com](https://learn.microsoft.com/en-us/azure/architecture/patterns/anti-corruption-layer?utm_source=chatgpt.com))。

## 价值与收益
- **模型一致性**：保持核心模型纯粹性，有助于长期演进和维护([ddd-practitioners.com](https://ddd-practitioners.com/home/glossary/bounded-context/bounded-context-relationship/anticorruption-layer/?utm_source=chatgpt.com), [codeopinion.com](https://codeopinion.com/anti-corruption-layer-for-mapping-between-boundaries/?utm_source=chatgpt.com))。
- **解耦合**：降低系统间直接依赖，使团队能够独立演进各自上下文([linkedin.com](https://www.linkedin.com/pulse/architecture-learnings-part-8-anti-corruption-layer-real-life-das?utm_source=chatgpt.com), [codeopinion.com](https://codeopinion.com/anti-corruption-layer-for-mapping-between-boundaries/?utm_source=chatgpt.com))。
- **可测试性**：通过中介层模拟和隔离外部依赖，提升单元测试和集成测试可靠性。

## 结论
防腐层是DDD架构中用于保护领域模型完整性和保持上下文清晰边界的重要模式。通过在不同系统或上下文之间引入专门的转换层，可以在保证系统演进速度的同时，维护模型的纯粹性和一致性，使软件更加健壮和易于维护。

## 直观示例
### 场景说明
假设我们有一个电商系统，需要与一个遗留的ERP系统（ExternalERP）交互，该系统中的订单模型与本地Order模型差异很大：

- **External ERP 的订单模型**：
  ```java
  class ExternalOrderDTO {
      String externalId;
      String customerCode;
      String productCodes; // CSV 格式
      double totalAmount;
  }
  ```
- **本地领域模型 Order**：
  ```java
  class Order {
      OrderId id;
      Customer customer;
      List<Product> products;
      Money amount;
  }
  ```

### 序列图（ASCII 示意）
```text
Customer  -> OrderService : createOrder(command)
OrderService -> ACLFacade    : sendToERP(mappedDTO)
ACLFacade  -> Adapter        : callExternalERP(dto)
Adapter    -> ExternalERP    : HTTP POST /orders
ExternalERP-> Adapter        : HTTP 200 OK (externalId)
Adapter    -> Translator     : toLocalOrder(resultDTO)
Translator -> ACLFacade      : localOrder
ACLFacade  -> OrderService   : orderConfirmation
OrderService-> Customer      : orderConfirmed(response)
```

### 代码示例
```java
// ACL Facade 层：简化对外部系统的调用
public class OrderACLFacade {
    private final ExternalOrderAdapter adapter;
    private final OrderTranslator translator;

    public OrderConfirmation createOrder(OrderCommand cmd) {
        ExternalOrderDTO dto = translator.toExternalDTO(cmd);
        ExternalOrderDTO result = adapter.sendOrder(dto);
        Order local = translator.toLocalOrder(result);
        // 保存本地领域模型...
        return new OrderConfirmation(local.getId());
    }
}

// Adapter：对接外部系统（如通过 REST API）
public class ExternalOrderAdapter {
    public ExternalOrderDTO sendOrder(ExternalOrderDTO dto) {
        // 调用 HTTP/消息队列等方式与 ERP 系统通信
        // ......
        return resultDTO;
    }
}

// Translator：在 DTO 和领域模型之间转换
public class OrderTranslator {
    public ExternalOrderDTO toExternalDTO(OrderCommand cmd) {
        // 将本地命令映射成外部系统格式
    }
    public Order toLocalOrder(ExternalOrderDTO dto) {
        // 将外部返回映射成本地领域模型
    }
}
```

以上示例通过 **ACLFacade**、**Adapter** 和 **Translator** 三层协作：

1. **Facade** 负责对外提供简洁接口，屏蔽底层细节。
2. **Adapter** 与外部系统直接通信，屏蔽协议与传输细节。
3. **Translator** 完成外部 DTO 与本地领域模型的双向映射。

这样一来，本地领域模型完全不依赖于外部系统的内部结构，既能直观感受到 ACL 的解耦效果，又能保证核心业务模型的纯粹性和可维护性。

