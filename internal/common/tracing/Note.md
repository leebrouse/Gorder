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
- 链路追踪部分在"./internal/common/trace"