package coreaudio

/*
#cgo darwin LDFLAGS: -framework AudioUnit -framework CoreAudio
#include <CoreAudio/AudioHardware.h>

*/
import "C"
import (
    "reflect"
    "unsafe"
)

const (
    AudioHardwareNoError = C.kAudioHardwareNoError;
    AudioHardwareNotRunningError = C.kAudioHardwareNotRunningError;
    AudioHardwareUnspecifiedError = C.kAudioHardwareUnspecifiedError;
    AudioHardwareUnknownPropertyError = C.kAudioHardwareUnknownPropertyError;
    AudioHardwareBadPropertySizeError = C.kAudioHardwareBadPropertySizeError;
    AudioHardwareIllegalOperationError = C.kAudioHardwareIllegalOperationError;
    AudioHardwareBadObjectError = C.kAudioHardwareBadObjectError;
    AudioHardwareBadDeviceError = C.kAudioHardwareBadDeviceError;
    AudioHardwareBadStreamError = C.kAudioHardwareBadStreamError;
    AudioHardwareUnsupportedOperationError = C.kAudioHardwareUnsupportedOperationError;
    AudioDeviceUnsupportedFormatError = C.kAudioDeviceUnsupportedFormatError;
    AudioDevicePermissionsError = C.kAudioDevicePermissionsError;
)

const (
    AudioObjectUnknown = C.kAudioObjectUnknown;
)

const (
    AudioObjectPropertyScopeGlobal = C.kAudioObjectPropertyScopeGlobal;
    AudioObjectPropertyScopeInput = C.kAudioObjectPropertyScopeInput;
    AudioObjectPropertyScopeOutput = C.kAudioObjectPropertyScopeOutput;
    AudioObjectPropertyScopePlayThrough = C.kAudioObjectPropertyScopePlayThrough;
    AudioObjectPropertyElementMaster = C.kAudioObjectPropertyElementMaster;
)

const (
    AudioObjectPropertySelectorWildcard = C.kAudioObjectPropertySelectorWildcard;
    AudioObjectPropertyScopeWildcard = C.kAudioObjectPropertyScopeWildcard;
    AudioObjectPropertyElementWildcard = C.kAudioObjectPropertyElementWildcard;
    AudioObjectClassIDWildcard = C.kAudioObjectClassIDWildcard;
)

const (
    AudioObjectClassID = C.kAudioObjectClassID;
)

const (
    AudioObjectPropertyBaseClass = C.kAudioObjectPropertyBaseClass;
    AudioObjectPropertyClass = C.kAudioObjectPropertyClass;
    AudioObjectPropertyOwner = C.kAudioObjectPropertyOwner;
    AudioObjectPropertyName = C.kAudioObjectPropertyName;
    AudioObjectPropertyModelName = C.kAudioObjectPropertyModelName;
    AudioObjectPropertyManufacturer = C.kAudioObjectPropertyManufacturer;
    AudioObjectPropertyElementName = C.kAudioObjectPropertyElementName;
    AudioObjectPropertyElementCategoryName = C.kAudioObjectPropertyElementCategoryName;
    AudioObjectPropertyElementNumberName = C.kAudioObjectPropertyElementNumberName;
    AudioObjectPropertyOwnedObjects = C.kAudioObjectPropertyOwnedObjects;
)

const (
    AudioPlugInClassID = C.kAudioPlugInClassID;
)

const (
    AudioPlugInPropertyBundleID = C.kAudioPlugInPropertyBundleID;
    AudioPlugInPropertyDeviceList = C.kAudioPlugInPropertyDeviceList;
    AudioPlugInPropertyTranslateUIDToDevice = C.kAudioPlugInPropertyTranslateUIDToDevice;
)

const (
    AudioTransportManagerClassID = C.kAudioTransportManagerClassID;
)

const (
    AudioTransportManagerPropertyEndPointList = C.kAudioTransportManagerPropertyEndPointList;
    AudioTransportManagerPropertyTranslateUIDToEndPoint = C.kAudioTransportManagerPropertyTranslateUIDToEndPoint;
    AudioTransportManagerPropertyTransportType = C.kAudioTransportManagerPropertyTransportType;
)

const (
    AudioDeviceClassID = C.kAudioDeviceClassID;
)

const (
    AudioDeviceTransportTypeUnknown = C.kAudioDeviceTransportTypeUnknown;
    AudioDeviceTransportTypeBuiltIn = C.kAudioDeviceTransportTypeBuiltIn;
    AudioDeviceTransportTypeAggregate = C.kAudioDeviceTransportTypeAggregate;
    AudioDeviceTransportTypeAutoAggregate = C.kAudioDeviceTransportTypeAutoAggregate;
    AudioDeviceTransportTypeVirtual = C.kAudioDeviceTransportTypeVirtual;
    AudioDeviceTransportTypePCI = C.kAudioDeviceTransportTypePCI;
    AudioDeviceTransportTypeUSB = C.kAudioDeviceTransportTypeUSB;
    AudioDeviceTransportTypeFireWire = C.kAudioDeviceTransportTypeFireWire;
    AudioDeviceTransportTypeBluetooth = C.kAudioDeviceTransportTypeBluetooth;
    AudioDeviceTransportTypeHDMI = C.kAudioDeviceTransportTypeHDMI;
    AudioDeviceTransportTypeDisplayPort = C.kAudioDeviceTransportTypeDisplayPort;
    AudioDeviceTransportTypeAirPlay = C.kAudioDeviceTransportTypeAirPlay;
    AudioDeviceTransportTypeAVB = C.kAudioDeviceTransportTypeAVB;
    AudioDeviceTransportTypeThunderbolt = C.kAudioDeviceTransportTypeThunderbolt;
)

const (
    AudioDevicePropertyConfigurationApplication = C.kAudioDevicePropertyConfigurationApplication;
    AudioDevicePropertyDeviceUID = C.kAudioDevicePropertyDeviceUID;
    AudioDevicePropertyModelUID = C.kAudioDevicePropertyModelUID;
    AudioDevicePropertyTransportType = C.kAudioDevicePropertyTransportType;
    AudioDevicePropertyRelatedDevices = C.kAudioDevicePropertyRelatedDevices;
    AudioDevicePropertyClockDomain = C.kAudioDevicePropertyClockDomain;
    AudioDevicePropertyDeviceIsAlive = C.kAudioDevicePropertyDeviceIsAlive;
    AudioDevicePropertyDeviceIsRunning = C.kAudioDevicePropertyDeviceIsRunning;
    AudioDevicePropertyDeviceCanBeDefaultDevice = C.kAudioDevicePropertyDeviceCanBeDefaultDevice;
    AudioDevicePropertyDeviceCanBeDefaultSystemDevice = C.kAudioDevicePropertyDeviceCanBeDefaultSystemDevice;
    AudioDevicePropertyLatency = C.kAudioDevicePropertyLatency;
    AudioDevicePropertyStreams = C.kAudioDevicePropertyStreams;
    AudioObjectPropertyControlList = C.kAudioObjectPropertyControlList;
    AudioDevicePropertySafetyOffset = C.kAudioDevicePropertySafetyOffset;
    AudioDevicePropertyNominalSampleRate = C.kAudioDevicePropertyNominalSampleRate;
    AudioDevicePropertyAvailableNominalSampleRates = C.kAudioDevicePropertyAvailableNominalSampleRates;
    AudioDevicePropertyIcon = C.kAudioDevicePropertyIcon;
    AudioDevicePropertyIsHidden = C.kAudioDevicePropertyIsHidden;
    AudioDevicePropertyPreferredChannelsForStereo = C.kAudioDevicePropertyPreferredChannelsForStereo;
    AudioDevicePropertyPreferredChannelLayout = C.kAudioDevicePropertyPreferredChannelLayout;
)

const (
    AudioEndPointDeviceClassID = C.kAudioEndPointDeviceClassID;
)

const (
    AudioEndPointDevicePropertyComposition = C.kAudioEndPointDevicePropertyComposition;
    AudioEndPointDevicePropertyEndPointList = C.kAudioEndPointDevicePropertyEndPointList;
    AudioEndPointDevicePropertyIsPrivate = C.kAudioEndPointDevicePropertyIsPrivate;
)

const (
    AudioEndPointClassID = C.kAudioEndPointClassID;
)

const (
    AudioStreamClassID = C.kAudioStreamClassID;
)

const (
    AudioStreamTerminalTypeUnknown = C.kAudioStreamTerminalTypeUnknown;
    AudioStreamTerminalTypeLine = C.kAudioStreamTerminalTypeLine;
    AudioStreamTerminalTypeDigitalAudioInterface = C.kAudioStreamTerminalTypeDigitalAudioInterface;
    AudioStreamTerminalTypeSpeaker = C.kAudioStreamTerminalTypeSpeaker;
    AudioStreamTerminalTypeHeadphones = C.kAudioStreamTerminalTypeHeadphones;
    AudioStreamTerminalTypeLFESpeaker = C.kAudioStreamTerminalTypeLFESpeaker;
    AudioStreamTerminalTypeReceiverSpeaker = C.kAudioStreamTerminalTypeReceiverSpeaker;
    AudioStreamTerminalTypeMicrophone = C.kAudioStreamTerminalTypeMicrophone;
    AudioStreamTerminalTypeHeadsetMicrophone = C.kAudioStreamTerminalTypeHeadsetMicrophone;
    AudioStreamTerminalTypeReceiverMicrophone = C.kAudioStreamTerminalTypeReceiverMicrophone;
    AudioStreamTerminalTypeTTY = C.kAudioStreamTerminalTypeTTY;
    AudioStreamTerminalTypeHDMI = C.kAudioStreamTerminalTypeHDMI;
    AudioStreamTerminalTypeDisplayPort = C.kAudioStreamTerminalTypeDisplayPort;
)

const (
    AudioStreamPropertyIsActive = C.kAudioStreamPropertyIsActive;
    AudioStreamPropertyDirection = C.kAudioStreamPropertyDirection;
    AudioStreamPropertyTerminalType = C.kAudioStreamPropertyTerminalType;
    AudioStreamPropertyStartingChannel = C.kAudioStreamPropertyStartingChannel;
    AudioStreamPropertyLatency = C.kAudioStreamPropertyLatency;
    AudioStreamPropertyVirtualFormat = C.kAudioStreamPropertyVirtualFormat;
    AudioStreamPropertyAvailableVirtualFormats = C.kAudioStreamPropertyAvailableVirtualFormats;
    AudioStreamPropertyPhysicalFormat = C.kAudioStreamPropertyPhysicalFormat;
    AudioStreamPropertyAvailablePhysicalFormats = C.kAudioStreamPropertyAvailablePhysicalFormats;
)

const (
    AudioControlClassID = C.kAudioControlClassID;
)

const (
    AudioControlPropertyScope = C.kAudioControlPropertyScope;
    AudioControlPropertyElement = C.kAudioControlPropertyElement;
)

const (
    AudioSliderControlClassID = C.kAudioSliderControlClassID;
)

const (
    AudioObjectSystemObject = C.kAudioObjectSystemObject;
)

const (
    AudioSliderControlPropertyValue = C.kAudioSliderControlPropertyValue;
    AudioSliderControlPropertyRange = C.kAudioSliderControlPropertyRange;
)

const (
    AudioLevelControlClassID = C.kAudioLevelControlClassID;
    AudioVolumeControlClassID = C.kAudioVolumeControlClassID;
    AudioLFEVolumeControlClassID = C.kAudioLFEVolumeControlClassID;
)

const (
    AudioLevelControlPropertyScalarValue = C.kAudioLevelControlPropertyScalarValue;
    AudioLevelControlPropertyDecibelValue = C.kAudioLevelControlPropertyDecibelValue;
    AudioLevelControlPropertyDecibelRange = C.kAudioLevelControlPropertyDecibelRange;
    AudioLevelControlPropertyConvertScalarToDecibels = C.kAudioLevelControlPropertyConvertScalarToDecibels;
    AudioLevelControlPropertyConvertDecibelsToScalar = C.kAudioLevelControlPropertyConvertDecibelsToScalar;
)

const (
    AudioBooleanControlClassID = C.kAudioBooleanControlClassID;
    AudioMuteControlClassID = C.kAudioMuteControlClassID;
    AudioSoloControlClassID = C.kAudioSoloControlClassID;
    AudioJackControlClassID = C.kAudioJackControlClassID;
    AudioLFEMuteControlClassID = C.kAudioLFEMuteControlClassID;
    AudioPhantomPowerControlClassID = C.kAudioPhantomPowerControlClassID;
    AudioPhaseInvertControlClassID = C.kAudioPhaseInvertControlClassID;
    AudioClipLightControlClassID = C.kAudioClipLightControlClassID;
    AudioTalkbackControlClassID = C.kAudioTalkbackControlClassID;
    AudioListenbackControlClassID = C.kAudioListenbackControlClassID;
)

const (
    AudioBooleanControlPropertyValue = C.kAudioBooleanControlPropertyValue;
)

const (
    AudioSelectorControlClassID = C.kAudioSelectorControlClassID;
    AudioDataSourceControlClassID = C.kAudioDataSourceControlClassID;
    AudioDataDestinationControlClassID = C.kAudioDataDestinationControlClassID;
    AudioClockSourceControlClassID = C.kAudioClockSourceControlClassID;
    AudioLineLevelControlClassID = C.kAudioLineLevelControlClassID;
    AudioHighPassFilterControlClassID = C.kAudioHighPassFilterControlClassID;
)

const (
    AudioSelectorControlPropertyCurrentItem = C.kAudioSelectorControlPropertyCurrentItem;
    AudioSelectorControlPropertyAvailableItems = C.kAudioSelectorControlPropertyAvailableItems;
    AudioSelectorControlPropertyItemName = C.kAudioSelectorControlPropertyItemName;
    AudioSelectorControlPropertyItemKind = C.kAudioSelectorControlPropertyItemKind;
)

const (
    AudioSelectorControlItemKindSpacer = C.kAudioSelectorControlItemKindSpacer;
)

const (
    AudioClockSourceItemKindInternal = C.kAudioClockSourceItemKindInternal;
)

const (
    AudioStereoPanControlClassID = C.kAudioStereoPanControlClassID;
)

const (
    AudioStereoPanControlPropertyValue = C.kAudioStereoPanControlPropertyValue;
    AudioStereoPanControlPropertyPanningChannels = C.kAudioStereoPanControlPropertyPanningChannels;
)

const (
    AudioObjectPropertyCreator = C.kAudioObjectPropertyCreator;
    AudioObjectPropertyListenerAdded = C.kAudioObjectPropertyListenerAdded;
    AudioObjectPropertyListenerRemoved = C.kAudioObjectPropertyListenerRemoved;
)

const (
    AudioSystemObjectClassID = C.kAudioSystemObjectClassID;
)

const (
    AudioHardwarePropertyDevices = C.kAudioHardwarePropertyDevices;
    AudioHardwarePropertyDefaultInputDevice = C.kAudioHardwarePropertyDefaultInputDevice;
    AudioHardwarePropertyDefaultOutputDevice = C.kAudioHardwarePropertyDefaultOutputDevice;
    AudioHardwarePropertyDefaultSystemOutputDevice = C.kAudioHardwarePropertyDefaultSystemOutputDevice;
    AudioHardwarePropertyTranslateUIDToDevice = C.kAudioHardwarePropertyTranslateUIDToDevice;
    AudioHardwarePropertyMixStereoToMono = C.kAudioHardwarePropertyMixStereoToMono;
    AudioHardwarePropertyTransportManagerList = C.kAudioHardwarePropertyTransportManagerList;
    AudioHardwarePropertyTranslateBundleIDToTransportManager = C.kAudioHardwarePropertyTranslateBundleIDToTransportManager;
    AudioHardwarePropertyProcessIsMaster = C.kAudioHardwarePropertyProcessIsMaster;
    AudioHardwarePropertyIsInitingOrExiting = C.kAudioHardwarePropertyIsInitingOrExiting;
    AudioHardwarePropertyUserIDChanged = C.kAudioHardwarePropertyUserIDChanged;
    AudioHardwarePropertyProcessIsAudible = C.kAudioHardwarePropertyProcessIsAudible;
    AudioHardwarePropertySleepingIsAllowed = C.kAudioHardwarePropertySleepingIsAllowed;
    AudioHardwarePropertyUnloadingIsAllowed = C.kAudioHardwarePropertyUnloadingIsAllowed;
    AudioHardwarePropertyHogModeIsAllowed = C.kAudioHardwarePropertyHogModeIsAllowed;
    AudioHardwarePropertyPlugInForBundleID = C.kAudioHardwarePropertyPlugInForBundleID;
    AudioHardwarePropertyUserSessionIsActiveOrHeadless = C.kAudioHardwarePropertyUserSessionIsActiveOrHeadless;
    AudioHardwarePropertyServiceRestarted = C.kAudioHardwarePropertyServiceRestarted;
)

const (
    AudioPlugInCreateAggregateDevice = C.kAudioPlugInCreateAggregateDevice;
    AudioPlugInDestroyAggregateDevice = C.kAudioPlugInDestroyAggregateDevice;
)

const (
    AudioTransportManagerCreateEndPointDevice = C.kAudioTransportManagerCreateEndPointDevice;
    AudioTransportManagerDestroyEndPointDevice = C.kAudioTransportManagerDestroyEndPointDevice;
)

const (
    AudioDeviceStartTimeIsInputFlag = C.kAudioDeviceStartTimeIsInputFlag;
    AudioDeviceStartTimeDontConsultDeviceFlag = C.kAudioDeviceStartTimeDontConsultDeviceFlag;
    AudioDeviceStartTimeDontConsultHALFlag = C.kAudioDeviceStartTimeDontConsultHALFlag;
)

const (
    AudioDevicePropertyPlugIn = C.kAudioDevicePropertyPlugIn;
    AudioDevicePropertyDeviceHasChanged = C.kAudioDevicePropertyDeviceHasChanged;
    AudioDevicePropertyDeviceIsRunningSomewhere = C.kAudioDevicePropertyDeviceIsRunningSomewhere;
    AudioDeviceProcessorOverload = C.kAudioDeviceProcessorOverload;
    AudioDevicePropertyHogMode = C.kAudioDevicePropertyHogMode;
    AudioDevicePropertyBufferFrameSize = C.kAudioDevicePropertyBufferFrameSize;
    AudioDevicePropertyBufferFrameSizeRange = C.kAudioDevicePropertyBufferFrameSizeRange;
    AudioDevicePropertyUsesVariableBufferFrameSizes = C.kAudioDevicePropertyUsesVariableBufferFrameSizes;
    AudioDevicePropertyIOCycleUsage = C.kAudioDevicePropertyIOCycleUsage;
    AudioDevicePropertyStreamConfiguration = C.kAudioDevicePropertyStreamConfiguration;
    AudioDevicePropertyIOProcStreamUsage = C.kAudioDevicePropertyIOProcStreamUsage;
    AudioDevicePropertyActualSampleRate = C.kAudioDevicePropertyActualSampleRate;
)

const (
    AudioDevicePropertyJackIsConnected = C.kAudioDevicePropertyJackIsConnected;
    AudioDevicePropertyVolumeScalar = C.kAudioDevicePropertyVolumeScalar;
    AudioDevicePropertyVolumeDecibels = C.kAudioDevicePropertyVolumeDecibels;
    AudioDevicePropertyVolumeRangeDecibels = C.kAudioDevicePropertyVolumeRangeDecibels;
    AudioDevicePropertyVolumeScalarToDecibels = C.kAudioDevicePropertyVolumeScalarToDecibels;
    AudioDevicePropertyVolumeDecibelsToScalar = C.kAudioDevicePropertyVolumeDecibelsToScalar;
    AudioDevicePropertyStereoPan = C.kAudioDevicePropertyStereoPan;
    AudioDevicePropertyStereoPanChannels = C.kAudioDevicePropertyStereoPanChannels;
    AudioDevicePropertyMute = C.kAudioDevicePropertyMute;
    AudioDevicePropertySolo = C.kAudioDevicePropertySolo;
    AudioDevicePropertyPhantomPower = C.kAudioDevicePropertyPhantomPower;
    AudioDevicePropertyPhaseInvert = C.kAudioDevicePropertyPhaseInvert;
    AudioDevicePropertyClipLight = C.kAudioDevicePropertyClipLight;
    AudioDevicePropertyTalkback = C.kAudioDevicePropertyTalkback;
    AudioDevicePropertyListenback = C.kAudioDevicePropertyListenback;
    AudioDevicePropertyDataSource = C.kAudioDevicePropertyDataSource;
    AudioDevicePropertyDataSources = C.kAudioDevicePropertyDataSources;
    AudioDevicePropertyDataSourceNameForIDCFString = C.kAudioDevicePropertyDataSourceNameForIDCFString;
    AudioDevicePropertyDataSourceKindForID = C.kAudioDevicePropertyDataSourceKindForID;
    AudioDevicePropertyClockSource = C.kAudioDevicePropertyClockSource;
    AudioDevicePropertyClockSources = C.kAudioDevicePropertyClockSources;
    AudioDevicePropertyClockSourceNameForIDCFString = C.kAudioDevicePropertyClockSourceNameForIDCFString;
    AudioDevicePropertyClockSourceKindForID = C.kAudioDevicePropertyClockSourceKindForID;
    AudioDevicePropertyPlayThru = C.kAudioDevicePropertyPlayThru;
    AudioDevicePropertyPlayThruSolo = C.kAudioDevicePropertyPlayThruSolo;
    AudioDevicePropertyPlayThruVolumeScalar = C.kAudioDevicePropertyPlayThruVolumeScalar;
    AudioDevicePropertyPlayThruVolumeDecibels = C.kAudioDevicePropertyPlayThruVolumeDecibels;
    AudioDevicePropertyPlayThruVolumeRangeDecibels = C.kAudioDevicePropertyPlayThruVolumeRangeDecibels;
    AudioDevicePropertyPlayThruVolumeScalarToDecibels = C.kAudioDevicePropertyPlayThruVolumeScalarToDecibels;
    AudioDevicePropertyPlayThruVolumeDecibelsToScalar = C.kAudioDevicePropertyPlayThruVolumeDecibelsToScalar;
    AudioDevicePropertyPlayThruStereoPan = C.kAudioDevicePropertyPlayThruStereoPan;
    AudioDevicePropertyPlayThruStereoPanChannels = C.kAudioDevicePropertyPlayThruStereoPanChannels;
    AudioDevicePropertyPlayThruDestination = C.kAudioDevicePropertyPlayThruDestination;
    AudioDevicePropertyPlayThruDestinations = C.kAudioDevicePropertyPlayThruDestinations;
    AudioDevicePropertyPlayThruDestinationNameForIDCFString = C.kAudioDevicePropertyPlayThruDestinationNameForIDCFString;
    AudioDevicePropertyChannelNominalLineLevel = C.kAudioDevicePropertyChannelNominalLineLevel;
    AudioDevicePropertyChannelNominalLineLevels = C.kAudioDevicePropertyChannelNominalLineLevels;
    AudioDevicePropertyChannelNominalLineLevelNameForIDCFString = C.kAudioDevicePropertyChannelNominalLineLevelNameForIDCFString;
    AudioDevicePropertyHighPassFilterSetting = C.kAudioDevicePropertyHighPassFilterSetting;
    AudioDevicePropertyHighPassFilterSettings = C.kAudioDevicePropertyHighPassFilterSettings;
    AudioDevicePropertyHighPassFilterSettingNameForIDCFString = C.kAudioDevicePropertyHighPassFilterSettingNameForIDCFString;
    AudioDevicePropertySubVolumeScalar = C.kAudioDevicePropertySubVolumeScalar;
    AudioDevicePropertySubVolumeDecibels = C.kAudioDevicePropertySubVolumeDecibels;
    AudioDevicePropertySubVolumeRangeDecibels = C.kAudioDevicePropertySubVolumeRangeDecibels;
    AudioDevicePropertySubVolumeScalarToDecibels = C.kAudioDevicePropertySubVolumeScalarToDecibels;
    AudioDevicePropertySubVolumeDecibelsToScalar = C.kAudioDevicePropertySubVolumeDecibelsToScalar;
    AudioDevicePropertySubMute = C.kAudioDevicePropertySubMute;
)

const (
    AudioAggregateDevicePropertyFullSubDeviceList = C.kAudioAggregateDevicePropertyFullSubDeviceList;
    AudioAggregateDevicePropertyActiveSubDeviceList = C.kAudioAggregateDevicePropertyActiveSubDeviceList;
    AudioAggregateDevicePropertyComposition = C.kAudioAggregateDevicePropertyComposition;
    AudioAggregateDevicePropertyMasterSubDevice = C.kAudioAggregateDevicePropertyMasterSubDevice;
)

const (
    AudioSubDeviceClassID = C.kAudioSubDeviceClassID;
)

const (
    AudioSubDeviceDriftCompensationMinQuality = C.kAudioSubDeviceDriftCompensationMinQuality;
    AudioSubDeviceDriftCompensationLowQuality = C.kAudioSubDeviceDriftCompensationLowQuality;
    AudioSubDeviceDriftCompensationMediumQuality = C.kAudioSubDeviceDriftCompensationMediumQuality;
    AudioSubDeviceDriftCompensationHighQuality = C.kAudioSubDeviceDriftCompensationHighQuality;
    AudioSubDeviceDriftCompensationMaxQuality = C.kAudioSubDeviceDriftCompensationMaxQuality;
)

const (
    AudioSubDevicePropertyExtraLatency = C.kAudioSubDevicePropertyExtraLatency;
    AudioSubDevicePropertyDriftCompensation = C.kAudioSubDevicePropertyDriftCompensation;
    AudioSubDevicePropertyDriftCompensationQuality = C.kAudioSubDevicePropertyDriftCompensationQuality;
)

type AudioObjectPropertyAddress struct {
    Selector uint32
    Scope uint32
    Element uint32
}

type AudioObjectID uint32

func build_AudioObjectPropertyAddress(addr AudioObjectPropertyAddress) *C.AudioObjectPropertyAddress {
    return &C.AudioObjectPropertyAddress {
        C.AudioObjectPropertySelector(addr.Selector),
        C.AudioObjectPropertyScope(addr.Scope),
        C.AudioObjectPropertyElement(addr.Element),
    }
}

func AudioObjectShow(objectID AudioObjectID) {
    C.AudioObjectShow(C.AudioObjectID(objectID))
}

func AudioObjectHasProperty(objectID AudioObjectID, addr AudioObjectPropertyAddress) bool {
    return int(C.AudioObjectHasProperty(C.AudioObjectID(objectID), build_AudioObjectPropertyAddress(addr))) != 0
}

func AudioObjectIsPropertySettable(objectID AudioObjectID, addr AudioObjectPropertyAddress) (bool, OSStatus) {
    var retval C.Boolean
    status := C.AudioObjectIsPropertySettable(
        C.AudioObjectID(objectID),
        build_AudioObjectPropertyAddress(addr),
        &retval,
    )
    return retval != 0, OSStatus(status)
}

func AudioObjectGetPropertyDataSize(objectID AudioObjectID, addr AudioObjectPropertyAddress, data interface {}) (uint32, OSStatus) {
    var retval C.UInt32
    qualifierRv := reflect.ValueOf(data)
    _qualifierSize := qualifierRv.Type().Elem().Size() * uintptr(qualifierRv.Len())
    // assert qualifierRv.Type().Elem().Size() == qualifierRv.Type().Elem().Align()
    _qualifier := qualifierRv.Index(0).UnsafeAddr()
    status := C.AudioObjectGetPropertyDataSize(
        C.AudioObjectID(objectID),
        build_AudioObjectPropertyAddress(addr),
        C.UInt32(_qualifierSize),
        unsafe.Pointer(_qualifier),
        &retval,
    )
    return uint32(retval), OSStatus(status)
}

func AudioObjectGetPropertyData(objectID AudioObjectID, addr AudioObjectPropertyAddress, qualifier interface {}) ([]byte, OSStatus) {
    var _qualifier unsafe.Pointer
    var _qualifierSize uintptr

    if qualifier != nil {
        qualifierRv := reflect.ValueOf(qualifier)
        _qualifierSize = qualifierRv.Type().Elem().Size() * uintptr(qualifierRv.Len())
        // assert qualifierRv.Type().Elem().Size() == qualifierRv.Type().Elem().Align()
        _qualifier = unsafe.Pointer(qualifierRv.Index(0).UnsafeAddr())
    } else {
        _qualifierSize = 0
        _qualifier = nil
    }

    _dataSize, _status := AudioObjectGetPropertyDataSize(objectID, addr, qualifier)
    if _status != 0 {
        return nil, _status
    }

    dataSize := C.UInt32(_dataSize)

    retval := make([]byte, uint32(dataSize))
    status := C.AudioObjectGetPropertyData(
        C.AudioObjectID(objectID),
        build_AudioObjectPropertyAddress(addr),
        C.UInt32(_qualifierSize),
        unsafe.Pointer(_qualifier),
        &dataSize,
        unsafe.Pointer(&retval[0]),
    )
    return retval, OSStatus(status)
}

func AudioObjectGetPropertyData_SystemObject(addr AudioObjectPropertyAddress) (AudioObjectID, OSStatus) {
    var data C.AudioObjectID
    var dataSize C.UInt32
    dataSize = C.UInt32(unsafe.Alignof(data))
    status := C.AudioObjectGetPropertyData(
        C.AudioObjectID(AudioObjectSystemObject),
        build_AudioObjectPropertyAddress(addr),
        C.UInt32(0),
        unsafe.Pointer(nil),
        &dataSize,
        unsafe.Pointer(&data),
    )
    return AudioObjectID(data), OSStatus(status)
}


func AudioObjectSetPropertyData(objectID AudioObjectID, addr AudioObjectPropertyAddress, qualifier interface {}, data interface {}) OSStatus {
    qualifierRv := reflect.ValueOf(qualifier)
    _qualifierSize := qualifierRv.Type().Elem().Size() * uintptr(qualifierRv.Len())
    // assert qualifierRv.Type().Elem().Size() == qualifierRv.Type().Elem().Align()
    _qualifier := qualifierRv.Index(0).UnsafeAddr()

    rv := reflect.ValueOf(data)
    rawDataSize := rv.Type().Elem().Size() * uintptr(rv.Len())
    // assert rv.Type().Elem().Size() == rv.Type().Elem().Align()
    rawData := rv.Index(0).UnsafeAddr()

    return OSStatus(C.AudioObjectSetPropertyData(
        C.AudioObjectID(objectID),
        build_AudioObjectPropertyAddress(addr),
        C.UInt32(_qualifierSize),
        unsafe.Pointer(_qualifier),
        C.UInt32(rawDataSize),
        unsafe.Pointer(rawData),
    ))
}


