package ice

type ObjectAdapter interface {
	GetName() string
	GetCommunicator() Communicator
	Activate()
	Hold()
	WaitForHold()
	Deactivate()
	WaitForDeactivate()
	IsDeactivated() bool
	Destroy()
	Add(servant Object, id Identity) ObjectPrx
	AddFacet(servant Object, id Identity, facet string) ObjectPrx
	AddWithUUID(servant Object) ObjectPrx
	AddFacetWithUUID(servant Object, facet string) ObjectPrx
	AddDefaultServant(servant Object, category string)
	Remove(id Identity) ObjectPrx
	RemoveFacet(id Identity, facet string) ObjectPrx
	//	removeAllFacets(id Identity) FacetMap
	RemoveDefaultServant(category string) ObjectPrx
	Find(id Identity) ObjectPrx
	FindFacet(id Identity, facet string) ObjectPrx
	//  findAllFacets(id Identity )FacetMap
	FindByProxy(proxy ObjectPrx) ObjectPrx
	//addServantLocator(locator ServantLocator , category string )
	// removeServantLocator(category string )ServantLocator
	// findServantLocator(category string )ServantLocator
	FindDefaultServant(category string) ObjectPrx
	CreateProxy(id Identity) ObjectPrx
	CreateDirectProxy(id Identity) ObjectPrx
	CreateIndirectProxy(id Identity) ObjectPrx
	//  setLocator(Locator* loc);
	// Locator* getLocator();
	RefreshPublishedEndpoints()
	//  EndpointSeq getEndpoints();
	//	EndpointSeq getPublishedEndpoints();
}
