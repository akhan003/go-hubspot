# PublicPropertyAssociationFilterBranchFiltersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | **string** |  | [default to "UNIFIED_EVENTS_IN_LIST"]
**Property** | **string** |  | 
**Operation** | [**PublicPropertyFilterOperation**](PublicPropertyFilterOperation.md) |  | 
**CoalescingRefineBy** | [**PublicEventAnalyticsFilterCoalescingRefineBy**](PublicEventAnalyticsFilterCoalescingRefineBy.md) |  | 
**Operator** | **string** |  | 
**ListId** | **int32** |  | 
**ToObjectTypeId** | Pointer to **string** |  | [optional] 
**AssociationTypeId** | **int32** |  | 
**AssociationCategory** | **string** |  | 
**ToObjectType** | Pointer to **string** |  | [optional] 
**PruningRefineBy** | [**PublicEventAnalyticsFilterCoalescingRefineBy**](PublicEventAnalyticsFilterCoalescingRefineBy.md) |  | 
**EnableTracking** | Pointer to **bool** |  | [optional] 
**PageUrl** | **string** |  | 
**CtaName** | **string** |  | 
**EventId** | **string** |  | 
**FormId** | Pointer to **string** |  | [optional] 
**PageId** | **string** |  | 
**EventTypeId** | **string** |  | 
**FilterLines** | [**[]PublicEventFilterMetadata**](PublicEventFilterMetadata.md) |  | 
**SubscriptionIds** | **[]int64** |  | 
**AcceptedStatuses** | **[]string** |  | 
**SubscriptionType** | **string** |  | 
**Channel** | **string** |  | 
**AcceptedOptStates** | **[]string** |  | 
**BusinessUnitId** | Pointer to **int64** |  | [optional] 
**CampaignId** | **string** |  | 
**SurveyId** | **string** |  | 
**SurveyQuestion** | **string** |  | 
**SurveyAnswerRowId** | Pointer to **string** |  | [optional] 
**SurveyAnswerColId** | Pointer to **string** |  | [optional] 
**ValueComparison** | [**PublicPropertyFilterOperation**](PublicPropertyFilterOperation.md) |  | 
**WebinarId** | Pointer to **string** |  | [optional] 
**Level** | **string** |  | 
**EmailId** | **int32** |  | 
**AppId** | **int32** |  | 
**ClickUrl** | Pointer to **string** |  | [optional] 
**PrivacyName** | **string** |  | 
**EntityType** | **string** |  | 
**SearchTermType** | **string** |  | 
**SearchTerms** | **[]string** |  | 
**AdNetwork** | **string** |  | 
**Metadata** | Pointer to [**PublicInListFilterMetadata**](PublicInListFilterMetadata.md) |  | [optional] 
**PropertyWithObjectId** | **string** |  | 
**ShouldAccept** | **bool** |  | 
**Source** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicPropertyAssociationFilterBranchFiltersInner

`func NewPublicPropertyAssociationFilterBranchFiltersInner(filterType string, property string, operation PublicPropertyFilterOperation, coalescingRefineBy PublicEventAnalyticsFilterCoalescingRefineBy, operator string, listId int32, associationTypeId int32, associationCategory string, pruningRefineBy PublicEventAnalyticsFilterCoalescingRefineBy, pageUrl string, ctaName string, eventId string, pageId string, eventTypeId string, filterLines []PublicEventFilterMetadata, subscriptionIds []int64, acceptedStatuses []string, subscriptionType string, channel string, acceptedOptStates []string, campaignId string, surveyId string, surveyQuestion string, valueComparison PublicPropertyFilterOperation, level string, emailId int32, appId int32, privacyName string, entityType string, searchTermType string, searchTerms []string, adNetwork string, propertyWithObjectId string, shouldAccept bool, ) *PublicPropertyAssociationFilterBranchFiltersInner`

NewPublicPropertyAssociationFilterBranchFiltersInner instantiates a new PublicPropertyAssociationFilterBranchFiltersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPropertyAssociationFilterBranchFiltersInnerWithDefaults

`func NewPublicPropertyAssociationFilterBranchFiltersInnerWithDefaults() *PublicPropertyAssociationFilterBranchFiltersInner`

NewPublicPropertyAssociationFilterBranchFiltersInnerWithDefaults instantiates a new PublicPropertyAssociationFilterBranchFiltersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetProperty

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetProperty(v string)`

SetProperty sets Property field to given value.


### GetOperation

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetOperation() PublicPropertyFilterOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetOperationOk() (*PublicPropertyFilterOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetOperation(v PublicPropertyFilterOperation)`

SetOperation sets Operation field to given value.


### GetCoalescingRefineBy

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetCoalescingRefineBy() PublicEventAnalyticsFilterCoalescingRefineBy`

GetCoalescingRefineBy returns the CoalescingRefineBy field if non-nil, zero value otherwise.

### GetCoalescingRefineByOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetCoalescingRefineByOk() (*PublicEventAnalyticsFilterCoalescingRefineBy, bool)`

GetCoalescingRefineByOk returns a tuple with the CoalescingRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoalescingRefineBy

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetCoalescingRefineBy(v PublicEventAnalyticsFilterCoalescingRefineBy)`

SetCoalescingRefineBy sets CoalescingRefineBy field to given value.


### GetOperator

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetListId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetListId() int32`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetListIdOk() (*int32, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetListId(v int32)`

SetListId sets ListId field to given value.


### GetToObjectTypeId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetToObjectTypeId() string`

GetToObjectTypeId returns the ToObjectTypeId field if non-nil, zero value otherwise.

### GetToObjectTypeIdOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetToObjectTypeIdOk() (*string, bool)`

GetToObjectTypeIdOk returns a tuple with the ToObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToObjectTypeId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetToObjectTypeId(v string)`

SetToObjectTypeId sets ToObjectTypeId field to given value.

### HasToObjectTypeId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) HasToObjectTypeId() bool`

HasToObjectTypeId returns a boolean if a field has been set.

### GetAssociationTypeId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetAssociationTypeId() int32`

GetAssociationTypeId returns the AssociationTypeId field if non-nil, zero value otherwise.

### GetAssociationTypeIdOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetAssociationTypeIdOk() (*int32, bool)`

GetAssociationTypeIdOk returns a tuple with the AssociationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationTypeId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetAssociationTypeId(v int32)`

SetAssociationTypeId sets AssociationTypeId field to given value.


### GetAssociationCategory

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetAssociationCategory() string`

GetAssociationCategory returns the AssociationCategory field if non-nil, zero value otherwise.

### GetAssociationCategoryOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetAssociationCategoryOk() (*string, bool)`

GetAssociationCategoryOk returns a tuple with the AssociationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCategory

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetAssociationCategory(v string)`

SetAssociationCategory sets AssociationCategory field to given value.


### GetToObjectType

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetToObjectType() string`

GetToObjectType returns the ToObjectType field if non-nil, zero value otherwise.

### GetToObjectTypeOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetToObjectTypeOk() (*string, bool)`

GetToObjectTypeOk returns a tuple with the ToObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToObjectType

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetToObjectType(v string)`

SetToObjectType sets ToObjectType field to given value.

### HasToObjectType

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) HasToObjectType() bool`

HasToObjectType returns a boolean if a field has been set.

### GetPruningRefineBy

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetPruningRefineBy() PublicEventAnalyticsFilterCoalescingRefineBy`

GetPruningRefineBy returns the PruningRefineBy field if non-nil, zero value otherwise.

### GetPruningRefineByOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetPruningRefineByOk() (*PublicEventAnalyticsFilterCoalescingRefineBy, bool)`

GetPruningRefineByOk returns a tuple with the PruningRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruningRefineBy

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetPruningRefineBy(v PublicEventAnalyticsFilterCoalescingRefineBy)`

SetPruningRefineBy sets PruningRefineBy field to given value.


### GetEnableTracking

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetEnableTracking() bool`

GetEnableTracking returns the EnableTracking field if non-nil, zero value otherwise.

### GetEnableTrackingOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetEnableTrackingOk() (*bool, bool)`

GetEnableTrackingOk returns a tuple with the EnableTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTracking

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetEnableTracking(v bool)`

SetEnableTracking sets EnableTracking field to given value.

### HasEnableTracking

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) HasEnableTracking() bool`

HasEnableTracking returns a boolean if a field has been set.

### GetPageUrl

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetPageUrl() string`

GetPageUrl returns the PageUrl field if non-nil, zero value otherwise.

### GetPageUrlOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetPageUrlOk() (*string, bool)`

GetPageUrlOk returns a tuple with the PageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageUrl

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetPageUrl(v string)`

SetPageUrl sets PageUrl field to given value.


### GetCtaName

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetCtaName() string`

GetCtaName returns the CtaName field if non-nil, zero value otherwise.

### GetCtaNameOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetCtaNameOk() (*string, bool)`

GetCtaNameOk returns a tuple with the CtaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtaName

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetCtaName(v string)`

SetCtaName sets CtaName field to given value.


### GetEventId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetFormId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetFormId() string`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetFormIdOk() (*string, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetFormId(v string)`

SetFormId sets FormId field to given value.

### HasFormId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) HasFormId() bool`

HasFormId returns a boolean if a field has been set.

### GetPageId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetPageId(v string)`

SetPageId sets PageId field to given value.


### GetEventTypeId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetEventTypeId() string`

GetEventTypeId returns the EventTypeId field if non-nil, zero value otherwise.

### GetEventTypeIdOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetEventTypeIdOk() (*string, bool)`

GetEventTypeIdOk returns a tuple with the EventTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetEventTypeId(v string)`

SetEventTypeId sets EventTypeId field to given value.


### GetFilterLines

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetFilterLines() []PublicEventFilterMetadata`

GetFilterLines returns the FilterLines field if non-nil, zero value otherwise.

### GetFilterLinesOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetFilterLinesOk() (*[]PublicEventFilterMetadata, bool)`

GetFilterLinesOk returns a tuple with the FilterLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterLines

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetFilterLines(v []PublicEventFilterMetadata)`

SetFilterLines sets FilterLines field to given value.


### GetSubscriptionIds

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetSubscriptionIds() []int64`

GetSubscriptionIds returns the SubscriptionIds field if non-nil, zero value otherwise.

### GetSubscriptionIdsOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetSubscriptionIdsOk() (*[]int64, bool)`

GetSubscriptionIdsOk returns a tuple with the SubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIds

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetSubscriptionIds(v []int64)`

SetSubscriptionIds sets SubscriptionIds field to given value.


### GetAcceptedStatuses

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetAcceptedStatuses() []string`

GetAcceptedStatuses returns the AcceptedStatuses field if non-nil, zero value otherwise.

### GetAcceptedStatusesOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetAcceptedStatusesOk() (*[]string, bool)`

GetAcceptedStatusesOk returns a tuple with the AcceptedStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedStatuses

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetAcceptedStatuses(v []string)`

SetAcceptedStatuses sets AcceptedStatuses field to given value.


### GetSubscriptionType

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetSubscriptionType() string`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetSubscriptionTypeOk() (*string, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetSubscriptionType(v string)`

SetSubscriptionType sets SubscriptionType field to given value.


### GetChannel

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetAcceptedOptStates

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetAcceptedOptStates() []string`

GetAcceptedOptStates returns the AcceptedOptStates field if non-nil, zero value otherwise.

### GetAcceptedOptStatesOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetAcceptedOptStatesOk() (*[]string, bool)`

GetAcceptedOptStatesOk returns a tuple with the AcceptedOptStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedOptStates

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetAcceptedOptStates(v []string)`

SetAcceptedOptStates sets AcceptedOptStates field to given value.


### GetBusinessUnitId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetBusinessUnitId() int64`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetBusinessUnitIdOk() (*int64, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetBusinessUnitId(v int64)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### GetCampaignId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.


### GetSurveyId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetSurveyId() string`

GetSurveyId returns the SurveyId field if non-nil, zero value otherwise.

### GetSurveyIdOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetSurveyIdOk() (*string, bool)`

GetSurveyIdOk returns a tuple with the SurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetSurveyId(v string)`

SetSurveyId sets SurveyId field to given value.


### GetSurveyQuestion

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetSurveyQuestion() string`

GetSurveyQuestion returns the SurveyQuestion field if non-nil, zero value otherwise.

### GetSurveyQuestionOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetSurveyQuestionOk() (*string, bool)`

GetSurveyQuestionOk returns a tuple with the SurveyQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyQuestion

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetSurveyQuestion(v string)`

SetSurveyQuestion sets SurveyQuestion field to given value.


### GetSurveyAnswerRowId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetSurveyAnswerRowId() string`

GetSurveyAnswerRowId returns the SurveyAnswerRowId field if non-nil, zero value otherwise.

### GetSurveyAnswerRowIdOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetSurveyAnswerRowIdOk() (*string, bool)`

GetSurveyAnswerRowIdOk returns a tuple with the SurveyAnswerRowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyAnswerRowId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetSurveyAnswerRowId(v string)`

SetSurveyAnswerRowId sets SurveyAnswerRowId field to given value.

### HasSurveyAnswerRowId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) HasSurveyAnswerRowId() bool`

HasSurveyAnswerRowId returns a boolean if a field has been set.

### GetSurveyAnswerColId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetSurveyAnswerColId() string`

GetSurveyAnswerColId returns the SurveyAnswerColId field if non-nil, zero value otherwise.

### GetSurveyAnswerColIdOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetSurveyAnswerColIdOk() (*string, bool)`

GetSurveyAnswerColIdOk returns a tuple with the SurveyAnswerColId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyAnswerColId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetSurveyAnswerColId(v string)`

SetSurveyAnswerColId sets SurveyAnswerColId field to given value.

### HasSurveyAnswerColId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) HasSurveyAnswerColId() bool`

HasSurveyAnswerColId returns a boolean if a field has been set.

### GetValueComparison

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetValueComparison() PublicPropertyFilterOperation`

GetValueComparison returns the ValueComparison field if non-nil, zero value otherwise.

### GetValueComparisonOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetValueComparisonOk() (*PublicPropertyFilterOperation, bool)`

GetValueComparisonOk returns a tuple with the ValueComparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueComparison

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetValueComparison(v PublicPropertyFilterOperation)`

SetValueComparison sets ValueComparison field to given value.


### GetWebinarId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetWebinarId() string`

GetWebinarId returns the WebinarId field if non-nil, zero value otherwise.

### GetWebinarIdOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetWebinarIdOk() (*string, bool)`

GetWebinarIdOk returns a tuple with the WebinarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebinarId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetWebinarId(v string)`

SetWebinarId sets WebinarId field to given value.

### HasWebinarId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) HasWebinarId() bool`

HasWebinarId returns a boolean if a field has been set.

### GetLevel

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetEmailId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetEmailId() int32`

GetEmailId returns the EmailId field if non-nil, zero value otherwise.

### GetEmailIdOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetEmailIdOk() (*int32, bool)`

GetEmailIdOk returns a tuple with the EmailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetEmailId(v int32)`

SetEmailId sets EmailId field to given value.


### GetAppId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### GetClickUrl

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetClickUrl() string`

GetClickUrl returns the ClickUrl field if non-nil, zero value otherwise.

### GetClickUrlOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetClickUrlOk() (*string, bool)`

GetClickUrlOk returns a tuple with the ClickUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickUrl

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetClickUrl(v string)`

SetClickUrl sets ClickUrl field to given value.

### HasClickUrl

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) HasClickUrl() bool`

HasClickUrl returns a boolean if a field has been set.

### GetPrivacyName

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetPrivacyName() string`

GetPrivacyName returns the PrivacyName field if non-nil, zero value otherwise.

### GetPrivacyNameOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetPrivacyNameOk() (*string, bool)`

GetPrivacyNameOk returns a tuple with the PrivacyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyName

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetPrivacyName(v string)`

SetPrivacyName sets PrivacyName field to given value.


### GetEntityType

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetSearchTermType

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetSearchTermType() string`

GetSearchTermType returns the SearchTermType field if non-nil, zero value otherwise.

### GetSearchTermTypeOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetSearchTermTypeOk() (*string, bool)`

GetSearchTermTypeOk returns a tuple with the SearchTermType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTermType

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetSearchTermType(v string)`

SetSearchTermType sets SearchTermType field to given value.


### GetSearchTerms

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetSearchTerms() []string`

GetSearchTerms returns the SearchTerms field if non-nil, zero value otherwise.

### GetSearchTermsOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetSearchTermsOk() (*[]string, bool)`

GetSearchTermsOk returns a tuple with the SearchTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerms

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetSearchTerms(v []string)`

SetSearchTerms sets SearchTerms field to given value.


### GetAdNetwork

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetAdNetwork() string`

GetAdNetwork returns the AdNetwork field if non-nil, zero value otherwise.

### GetAdNetworkOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetAdNetworkOk() (*string, bool)`

GetAdNetworkOk returns a tuple with the AdNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdNetwork

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetAdNetwork(v string)`

SetAdNetwork sets AdNetwork field to given value.


### GetMetadata

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetMetadata() PublicInListFilterMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetMetadataOk() (*PublicInListFilterMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetMetadata(v PublicInListFilterMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPropertyWithObjectId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetPropertyWithObjectId() string`

GetPropertyWithObjectId returns the PropertyWithObjectId field if non-nil, zero value otherwise.

### GetPropertyWithObjectIdOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetPropertyWithObjectIdOk() (*string, bool)`

GetPropertyWithObjectIdOk returns a tuple with the PropertyWithObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyWithObjectId

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetPropertyWithObjectId(v string)`

SetPropertyWithObjectId sets PropertyWithObjectId field to given value.


### GetShouldAccept

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetShouldAccept() bool`

GetShouldAccept returns the ShouldAccept field if non-nil, zero value otherwise.

### GetShouldAcceptOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetShouldAcceptOk() (*bool, bool)`

GetShouldAcceptOk returns a tuple with the ShouldAccept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldAccept

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetShouldAccept(v bool)`

SetShouldAccept sets ShouldAccept field to given value.


### GetSource

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PublicPropertyAssociationFilterBranchFiltersInner) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


