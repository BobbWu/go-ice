package ice

type Communicator interface {
	Destroy()
	Shutdown()
	WaitForShutdown()
	IsShutdown() bool
	StringToProxy(str string) ObjectPrx
	ProxyToString(obj ObjectPrx) string
	PropertyToProxy(property string) ObjectPrx
	// proxyToProperty(proxy ObjectPrx, property string ) PropertyDict
	StringToIdentity(str string) Identity
	IdentityToString(ident Identity) string
	CreateObjectAdapter(name string) ObjectAdapter
	CreateObjectAdapterWithEndpoints(name string, endpoints string) ObjectAdapter
	//createObjectAdapterWithRouter(name string , rtr Router* ) ObjectAdapter
	//addObjectFactory(factory ObjectFactory , id string )
	// findObjectFactory(id string )ObjectFactory
	//getImplicitContext()ImplicitContext
	// getProperties()Properties
	GetLogger() Logger
	// getObserver() Ice::Instrumentation::CommunicatorObserver
	//getDefaultRouter() Router*
	// setDefaultRouter(rtr Router* )
	// getDefaultLocator()  Locator*
	// setDefaultLocator(loc Locator* )
	//  getPluginManager() PluginManager
	FlushBatchRequests()
	CreateAdmin(adminAdapter ObjectAdapter, adminId Identity) ObjectPrx
	GetAdmin() ObjectPrx
	AddAdminFacet(servant Object, facet string)
	RemoveAdminFacet(facet string) ObjectPrx
	FindAdminFacet(facet string) ObjectPrx
	//findAllAdminFacets() FacetMap
}
