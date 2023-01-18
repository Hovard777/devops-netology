> Найдите, где перечислены все доступные ```resource``` и ```data_source```, приложите ссылку на эти строки в коде на гитхабе.

List DataSources https://github.com/hashicorp/terraform-provider-aws/blob/ef5f62e24b84a26c4f33031216424e739fe06b5b/internal/provider/provider.go#L419  

List of Resources https://github.com/hashicorp/terraform-provider-aws/blob/ef5f62e24b84a26c4f33031216424e739fe06b5b/internal/provider/provider.go#L944
    


>Для создания очереди сообщений SQS используется ресурс aws_sqs_queue у которого есть параметр name.  
>1. С каким другим параметром конфликтует ```name```? Приложите строчку кода, в которой это указано.  
>2. Какая максимальная длина имени?  
>3. Какому регулярному выражению должно подчиняться имя?  

1. ConflictsWith: []string{"name_prefix"} - https://github.com/hashicorp/terraform-provider-aws/blob/ef5f62e24b84a26c4f33031216424e739fe06b5b/internal/service/sqs/queue.go#L88  
2. 
3. 