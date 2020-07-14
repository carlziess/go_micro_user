## 使用go-micro将protobuf生成的服务给注册到注册中心（使用的是consul）
## gin框架加go-micro做http的API，实现路由的访问，从consul中拿取具体service
- 使用go-micro中装饰器的组件：wrapper
- 使用"github.com/afex/hystrix-go/hystrix"做熔断和降级