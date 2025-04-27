### Q:链路追踪中什么是export? 什么是 trace? 什么是span?
#### 名词解释:
1. Trace（追踪）	一个完整的请求在系统中的调用链条，代表“一次请求的全貌”。
2. Span（跨度）	是 Trace 的基本组成单元，代表“某个服务内某次操作”，比如一个函数调用或一次 HTTP 请求。
3. Export（导出）	将采集到的 Trace 数据导出到外部系统（如 Jaeger、Zipkin、Skywalking 或 OTLP Collector），用于可视化、分析或存储
#### Struct graph:
```
Trace（请求 ID: trace-id）
│
├── Span 1: 网关服务接收请求（根 Span）
│     ├── attributes: HTTP method, path, etc.
│     └── time: 10ms
│
├── Span 2: 服务 A 处理请求
│     ├── attributes: DB query, etc.
│     └── time: 30ms
│
├── Span 3: 服务 B 处理请求（被 A 调用）
│     ├── attributes: Cache miss, etc.
│     └── time: 15ms
│
└── Span 4: 服务 C 处理（被 B 调用）
├── attributes: External API call
└── time: 20ms

Export（导出）→ Jaeger / Zipkin / Skywalking / OpenTelemetry Collector
```

### Trace and Extract:

---
inject 和 extract 是 OpenTelemetry 中分布式上下文传播（context propagation）的核心机制，用于在服务之间传递 Trace 信息，从而实现完整的调用链跟踪。

✅ Inject 的作用：
将当前上下文中的 Trace 信息写入到消息或请求的载体（carrier）中，比如：

HTTP 请求头（Headers）

gRPC Metadata

RabbitMQ headers（你代码中的场景）

👉 通俗理解：

把“我是谁（TraceID, SpanID, 上下文）”这张小纸条，贴到即将发送的消息上，让下一个服务知道“我是从哪里来的”。

✅ Extract 的作用：
从接收到的消息或请求中提取 Trace 信息，恢复上下文，以便继续追踪：

如果消息中带有 traceparent、baggage、b3 等字段，它会把它们解析出来放进新的 context.Context 中。

后续创建的新 span 就自动成为“父 Trace”下的子节点。

👉 通俗理解：

拿到别人贴的小纸条，知道“这个请求来自哪条链路”，然后自己继续在这张链上添一笔。

🧠 一句话总结：
Inject 是写 trace到“信封”，
Extract 是从信封里读 trace出来。
二者共同完成链路信息在不同服务间的无缝传递，是分布式链路追踪的关键。
---

- 链路追踪部分在"./internal/common/trace"