## sink的redis缓存设计

1. 表映射的情况下，映射关系应当被缓存到redis，当rulemanager服务重启时加载映射关系。
2. 对于sink的缓存：
   1. endpoints
   2. user
   3. password
    使用以上配置拉生成一个sink
3. 对于表映射的缓存
   1. 缓存表字段及其映射关系

4. 复原：
   1. 使用sink配置生成sink及其tables
      1. sink不存在时，给出warning，丢弃sink及其映射关系
      2. 当sink存在
   2. 使用表映射缓存恢复映射关系
      1. 






