package ice

type OperationMode int

type ObjectPrx interface {
	Ice_GetCommunicator() Communicator

	Ice_IsA(id string) bool
	Ice_Ping()
	Ice_Ids() []string
	Ice_Id() string
	Ice_Invoke(operation string, mode OperationMode, inEncaps []byte, outEncaps []byte)
	Ice_GetIdentity() Identity
	Ice_Identity(newIdentity Identity) Identity
	Ice_GetContext() map[string]string
	Ice_Context(newContext map[string]string) ObjectPrx
	Ice_GetFacet() string
	Ice_Facet(newFacet string) ObjectPrx
	Ice_GetAdapterId() string
	Ice_AdapterId(newAdapterId string) ObjectPrx
	Ice_GetEndpoints() []Endpoint
	Ice_Endpoints(newEndPoints []Endpoint) ObjectPrx
	Ice_GetLocatorCacheTimeout() int
	Ice_LocatorCacheTimeout(timeout int) ObjectPrx
	Ice_GetInvocationTimeout() int
	Ice_InvocationTimeout(timeout int) ObjectPrx
	Ice_IsConnectionCached() bool
	Ice_ConnectionCached(newCache bool) ObjectPrx
	//	Ice_GetEndpointSelection() EndpointSelectionType
	//	Ice_EndpointSelection(newType EndpointSelectionType) ObjectPrx
	ice_IsSecure() bool
	ice_Secure(b bool) ObjectPrx
	//Ice_EncodingVersion(e Ice.EncodingVersion )ObjectPrx
	//Ice_GetEncodingVersion() Ice.EncodingVersion
	ice_IsPreferSecure() bool
	ice_PreferSecure(b bool) ObjectPrx
	// Ice_GetRouter()Ice.RouterPrx
	// Ice_Router(router Ice.RouterPrx )ObjectPrx
	// Ice_GetLocator()Ice.LocatorPrx
	// Ice_Locator(locator Ice.LocatorPrx )ObjectPrx
	Ice_IsCollocationOptimized() bool
	Ice_CollocationOptimized(b bool) ObjectPrx
	Ice_Twoway() ObjectPrx
	Ice_IsTwoway() bool
	Ice_Oneway() ObjectPrx
	Ice_IsOneway() bool
	Ice_BatchOneway() ObjectPrx
	Ice_IsBatchOneway() bool
	Ice_Datagram() ObjectPrx
	Ice_IsDatagram() bool
	Ice_BatchDatagram() ObjectPrx
	Ice_IsBatchDatagram() bool
	Ice_Compress(co bool) ObjectPrx
	Ice_Timeout(t int) ObjectPrx
	Ice_ConnectionId(connectionId string) ObjectPrx
	Ice_GetConnectionId() string
	Ice_GetConnection() Connection
	Ice_GetCachedConnection() Connection
	Ice_FlushBatchRequests()
}
