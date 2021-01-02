package topic

//特殊topic
const RpcResponseTopic = "/topic/vi/rpc/response"

//消息传递类topic
//  /topic/{version}/msg/{mainResource}/{subResource1}/{subResource2}
// version: v1
// resourceName: ant,task,networkDataFetcher...
// 当type为msg时，最后一项是entityName，代表一个实体名称，一般会对应一个具体的类型，比如task, networkDataFetch,networkDataFetcherConf等等
// 当type为rpc时，最后一项是action，代表一个动作，
// action: data类list, update, create, command类，stop， start， pause

const NetworkDataFetcherTopicPrefix = "/topic/v1/msg/networkDataFetcher"
const NetworkDataFetcherConfSetTopic = NetworkDataFetcherTopicPrefix + "/conf"
const NetworkDataFetcherStatusSetTopic = NetworkDataFetcherTopicPrefix + "/status"

const TaskTopicStartFix = "/topic/v1/msg/task"
const TaskStatusSetTopic = TaskTopicStartFix + "/status"
const TaskResultReportTopic = TaskTopicStartFix + "/telemetry/"

const QueenMetaTopic = "/topic/v1/data/queen"
const AntMetaTopic = "/topic/v1/data/ant"

//rpc类topic
//  /topic/{version}/rpc/{mainResource}/{subResource1}/{subResource2}/{action}
// version: v1
// resourceName: ant,task,networkDataFetcher...
// 当type为msg时，最后一项是entityName，代表一个实体名称，一般会对应一个具体的类型，比如task, networkDataFetch,networkDataFetcherConf等等
// 当type为rpc时，最后一项是action，代表一个动作，
// action: data类list, update, create, command类，stop， start， pause
