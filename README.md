# Rule-manager

规则引擎相关服务。设置规则将数据转发置你想要的目标。

## 使用
### 环境变量
设置系统环境变量：
```bash
# root@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local
export DSN="<your-dsn>"

export RuleTopic="<your-rule-topic>"
```
`DSN` 信息用于连接服务数据库，目前使用的是 *MySQL* 驱动。采用的是 GORM 所以可以适配多个不同数据驱动。

`RuleTopic` 信息用于创建订阅 ID, 订阅服务可解析此 ID 创建规则主题然后用于用户订阅数据传输。
## 依赖库
Core 的调用以及 Service 中的用户认证 Auth ，还有分页工具用的是 core-broker 项目中的包。后面注意拆解。
```go
import (
    "github.com/tkeel-io/core-broker/pkg/auth"
    "github.com/tkeel-io/core-broker/pkg/core"
    "github.com/tkeel-io/core-broker/pkg/deviceutil"
    "github.com/tkeel-io/core-broker/pkg/pagination"
)

```
## 设计
### 表设计
#### rules
规则表，用于存储规则信息。

| 字段 | 类型 | 备注                    |
| ---- | --- |-----------------------|
| id | int | 主键                    |
|user_id| string| 用户ID                  |
|sub_id|int| 订阅ID                  |
|sub_enpoint|string| 订阅地址                  |
|name|string| 规则名称                  |
|status|int| 规则状态: 0-未启动，1-启动      |
|desc|string| 规则描述                  |
|type|int| 规则类型: 0-消息，1-时序...可拓展 |
|created_at|DateTime| 创建时间                  |
|updated_at|DateTime| 更新时间                  |
|deleted_at|DateTime| 删除时间                  |

### rule_entities
用于存储规则与实体设备关系的中间表

| 字段 | 类型  | 备注                    |
| ---- |-----|-----------------------|
|unique_key| string  | 采用规则拼接的唯一键，主要用于避免重复插入 |
|rule_id| int | 规则ID                  |
|entity_id| string | 实体设备ID                  |

### targets
规则目标表，用于存储规则转发目标的一些信息。

| 字段 | 类型  | 备注                                             |
| ---- |-----|------------------------------------------------|
|id| int | 主键                                             |
|type| int | 目标类型: 1-kafka，2-对象存储...可拓展                     |
|host| string | 目标地址                                           |
|value| string | 目标值，比如说 kafka 的某一个主题，存储对象的某一个桶... 自定义而适配于自己的服务 |
|ext |json| 扩展信息，以json格式存储                                 |
|rule_id | int | 规则ID                                           |