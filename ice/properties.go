package ice

type Properties interface {
	GetProperty(key string) string
	GetPropertyWithDefault(key string, value string) string
	GetPropertyAsInt(key string) int
	GetPropertyAsIntWithDefault(key string, value int) int
	//getPropertyAsList(key string) StringSeq
	//getPropertyAsListWithDefault(key string, value StringSeq) StringSeq
	//getPropertiesForPrefix(prefix string) PropertyDict
	SetProperty(key string, value string)
	//getCommandLineOptions() StringSeq
	//parseCommandLineOptions(prefix string, options StringSeq) StringSeq
	//	parseIceCommandLineOptions(options StringSeq) StringSeq
	Load(file string)
	Clone() Properties
}
