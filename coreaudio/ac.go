package coreaudio

/*
#cgo darwin LDFLAGS: -framework AudioUnit
#include <AudioUnit/AudioUnit.h>
#include <CoreAudio/AudioHardware.h>

*/
import "C"
import (
    "reflect"
    "unsafe"
)

const (
    AudioUnitType_Output = C.kAudioUnitType_Output;
    AudioUnitType_MusicDevice = C.kAudioUnitType_MusicDevice;
    AudioUnitType_MusicEffect = C.kAudioUnitType_MusicEffect;
    AudioUnitType_FormatConverter = C.kAudioUnitType_FormatConverter;
    AudioUnitType_Effect = C.kAudioUnitType_Effect;
    AudioUnitType_Mixer = C.kAudioUnitType_Mixer;
    AudioUnitType_Panner = C.kAudioUnitType_Panner;
    AudioUnitType_Generator = C.kAudioUnitType_Generator;
    AudioUnitType_OfflineEffect = C.kAudioUnitType_OfflineEffect;
)
const (
	AudioUnitManufacturer_Apple = C.kAudioUnitManufacturer_Apple;
)
const (
    AudioUnitSubType_GenericOutput = C.kAudioUnitSubType_GenericOutput;
    AudioUnitSubType_HALOutput = C.kAudioUnitSubType_HALOutput;
    AudioUnitSubType_DefaultOutput = C.kAudioUnitSubType_DefaultOutput;
    AudioUnitSubType_SystemOutput = C.kAudioUnitSubType_SystemOutput;
    // AudioUnitSubType_RemoteIO = C.kAudioUnitSubType_RemoteIO;
    AudioUnitSubType_VoiceProcessingIO = C.kAudioUnitSubType_VoiceProcessingIO;
)
const (
    AudioUnitSubType_DLSSynth = C.kAudioUnitSubType_DLSSynth;
    AudioUnitSubType_Sampler = C.kAudioUnitSubType_Sampler;
)
const (
    AudioUnitSubType_AUConverter = C.kAudioUnitSubType_AUConverter;
    AudioUnitSubType_NewTimePitch = C.kAudioUnitSubType_NewTimePitch;
    AudioUnitSubType_TimePitch = C.kAudioUnitSubType_TimePitch;
    AudioUnitSubType_DeferredRenderer = C.kAudioUnitSubType_DeferredRenderer;
    AudioUnitSubType_Splitter = C.kAudioUnitSubType_Splitter;
    AudioUnitSubType_Merger = C.kAudioUnitSubType_Merger;
    AudioUnitSubType_Varispeed = C.kAudioUnitSubType_Varispeed;
    AudioUnitSubType_RoundTripAAC = C.kAudioUnitSubType_RoundTripAAC;
    // AudioUnitSubType_AUiPodTime = C.kAudioUnitSubType_AUiPodTime;
    // AudioUnitSubType_AUiPodTimeOther = C.kAudioUnitSubType_AUiPodTimeOther;
)

const (
    AudioUnitSubType_PeakLimiter = C.kAudioUnitSubType_PeakLimiter;
    AudioUnitSubType_DynamicsProcessor = C.kAudioUnitSubType_DynamicsProcessor;
    AudioUnitSubType_LowPassFilter = C.kAudioUnitSubType_LowPassFilter;
    AudioUnitSubType_HighPassFilter = C.kAudioUnitSubType_HighPassFilter;
    AudioUnitSubType_BandPassFilter = C.kAudioUnitSubType_BandPassFilter;
    AudioUnitSubType_HighShelfFilter = C.kAudioUnitSubType_HighShelfFilter;
    AudioUnitSubType_LowShelfFilter = C.kAudioUnitSubType_LowShelfFilter;
    AudioUnitSubType_ParametricEQ = C.kAudioUnitSubType_ParametricEQ;
    AudioUnitSubType_Distortion = C.kAudioUnitSubType_Distortion;
    AudioUnitSubType_Delay = C.kAudioUnitSubType_Delay;
    AudioUnitSubType_GraphicEQ = C.kAudioUnitSubType_GraphicEQ;
    AudioUnitSubType_MultiBandCompressor = C.kAudioUnitSubType_MultiBandCompressor;
    AudioUnitSubType_MatrixReverb = C.kAudioUnitSubType_MatrixReverb;
    AudioUnitSubType_SampleDelay = C.kAudioUnitSubType_SampleDelay;
    AudioUnitSubType_Pitch = C.kAudioUnitSubType_Pitch;
    AudioUnitSubType_AUFilter = C.kAudioUnitSubType_AUFilter;
    AudioUnitSubType_NetSend = C.kAudioUnitSubType_NetSend;
    AudioUnitSubType_RogerBeep = C.kAudioUnitSubType_RogerBeep;
    // AudioUnitSubType_Reverb2 = C.kAudioUnitSubType_Reverb2;
    // AudioUnitSubType_AUiPodEQ = C.kAudioUnitSubType_AUiPodEQ;
    // AudioUnitSubType_NBandEQ = C.kAudioUnitSubType_NBandEQ;
)
const (
    AudioUnitSubType_MultiChannelMixer = C.kAudioUnitSubType_MultiChannelMixer;
    AudioUnitSubType_StereoMixer = C.kAudioUnitSubType_StereoMixer;
    AudioUnitSubType_3DMixer = C.kAudioUnitSubType_3DMixer;
    AudioUnitSubType_MatrixMixer = C.kAudioUnitSubType_MatrixMixer;
    // AudioUnitSubType_AU3DMixerEmbedded = C.kAudioUnitSubType_AU3DMixerEmbedded;
)
const (
    ParameterEvent_Immediate = C.kParameterEvent_Immediate;
    ParameterEvent_Ramped = C.kParameterEvent_Ramped;
)
const (
    AudioUnitRange = C.kAudioUnitRange;
    AudioUnitInitializeSelect = C.kAudioUnitInitializeSelect;
    AudioUnitUninitializeSelect = C.kAudioUnitUninitializeSelect;
    AudioUnitGetPropertyInfoSelect = C.kAudioUnitGetPropertyInfoSelect;
    AudioUnitGetPropertySelect = C.kAudioUnitGetPropertySelect;
    AudioUnitSetPropertySelect = C.kAudioUnitSetPropertySelect;
    AudioUnitAddPropertyListenerSelect = C.kAudioUnitAddPropertyListenerSelect;
    AudioUnitRemovePropertyListenerSelect = C.kAudioUnitRemovePropertyListenerSelect;
    AudioUnitRemovePropertyListenerWithUserDataSelect = C.kAudioUnitRemovePropertyListenerWithUserDataSelect;
    AudioUnitAddRenderNotifySelect = C.kAudioUnitAddRenderNotifySelect;
    AudioUnitRemoveRenderNotifySelect = C.kAudioUnitRemoveRenderNotifySelect;
    AudioUnitGetParameterSelect = C.kAudioUnitGetParameterSelect;
    AudioUnitSetParameterSelect = C.kAudioUnitSetParameterSelect;
    AudioUnitScheduleParametersSelect = C.kAudioUnitScheduleParametersSelect;
    AudioUnitRenderSelect = C.kAudioUnitRenderSelect;
    AudioUnitResetSelect = C.kAudioUnitResetSelect;
    AudioUnitComplexRenderSelect = C.kAudioUnitComplexRenderSelect;
    AudioUnitProcessSelect = C.kAudioUnitProcessSelect;
    AudioUnitProcessMultipleSelect = C.kAudioUnitProcessMultipleSelect;
)
const (
    AudioUnitSubType_SphericalHeadPanner = C.kAudioUnitSubType_SphericalHeadPanner;
    AudioUnitSubType_VectorPanner = C.kAudioUnitSubType_VectorPanner;
    AudioUnitSubType_SoundFieldPanner = C.kAudioUnitSubType_SoundFieldPanner;
    AudioUnitSubType_HRTFPanner = C.kAudioUnitSubType_HRTFPanner;
)
const (
    AudioUnitSubType_NetReceive = C.kAudioUnitSubType_NetReceive;
    AudioUnitSubType_ScheduledSoundPlayer = C.kAudioUnitSubType_ScheduledSoundPlayer;
    AudioUnitSubType_AudioFilePlayer = C.kAudioUnitSubType_AudioFilePlayer;
)
const (
    AudioUnitRenderAction_PreRender = C.kAudioUnitRenderAction_PreRender;
    AudioUnitRenderAction_PostRender = C.kAudioUnitRenderAction_PostRender;
    AudioUnitRenderAction_OutputIsSilence = C.kAudioUnitRenderAction_OutputIsSilence;
    AudioOfflineUnitRenderAction_Preflight = C.kAudioOfflineUnitRenderAction_Preflight;
    AudioOfflineUnitRenderAction_Render = C.kAudioOfflineUnitRenderAction_Render;
    AudioOfflineUnitRenderAction_Complete = C.kAudioOfflineUnitRenderAction_Complete;
    AudioUnitRenderAction_PostRenderError = C.kAudioUnitRenderAction_PostRenderError;
    AudioUnitRenderAction_DoNotCheckRenderArgs = C.kAudioUnitRenderAction_DoNotCheckRenderArgs;
)
const (
    NoErr = C.noErr;
    AudioUnitErr_InvalidProperty = C.kAudioUnitErr_InvalidProperty;
    AudioUnitErr_InvalidParameter = C.kAudioUnitErr_InvalidParameter;
    AudioUnitErr_InvalidElement = C.kAudioUnitErr_InvalidElement;
    AudioUnitErr_NoConnection = C.kAudioUnitErr_NoConnection;
    AudioUnitErr_FailedInitialization = C.kAudioUnitErr_FailedInitialization;
    AudioUnitErr_TooManyFramesToProcess = C.kAudioUnitErr_TooManyFramesToProcess;
    AudioUnitErr_InvalidFile = C.kAudioUnitErr_InvalidFile;
    AudioUnitErr_FormatNotSupported = C.kAudioUnitErr_FormatNotSupported;
    AudioUnitErr_Uninitialized = C.kAudioUnitErr_Uninitialized;
    AudioUnitErr_InvalidScope = C.kAudioUnitErr_InvalidScope;
    AudioUnitErr_PropertyNotWritable = C.kAudioUnitErr_PropertyNotWritable;
    AudioUnitErr_CannotDoInCurrentContext = C.kAudioUnitErr_CannotDoInCurrentContext;
    AudioUnitErr_InvalidPropertyValue = C.kAudioUnitErr_InvalidPropertyValue;
    AudioUnitErr_PropertyNotInUse = C.kAudioUnitErr_PropertyNotInUse;
    AudioUnitErr_Initialized = C.kAudioUnitErr_Initialized;
    AudioUnitErr_InvalidOfflineRender = C.kAudioUnitErr_InvalidOfflineRender;
    AudioUnitErr_Unauthorized = C.kAudioUnitErr_Unauthorized;
    AudioUnitErr_IllegalInstrument = C.kAudioUnitErr_IllegalInstrument;
    AudioUnitErr_InstrumentTypeNotFound = C.kAudioUnitErr_InstrumentTypeNotFound;
    AudioUnitErr_UnknownFileType = C.kAudioUnitErr_UnknownFileType;
    AudioUnitErr_FileNotSpecified = C.kAudioUnitErr_FileNotSpecified;
)
const (
    AudioOutputUnitRange = C.kAudioOutputUnitRange;
    AudioOutputUnitStartSelect = C.kAudioOutputUnitStartSelect;
    AudioOutputUnitStopSelect = C.kAudioOutputUnitStopSelect;
)
const (
    AudioCodecPropertySupportedInputFormats = C.kAudioCodecPropertySupportedInputFormats;
    AudioCodecPropertySupportedOutputFormats = C.kAudioCodecPropertySupportedOutputFormats;
    AudioCodecPropertyAvailableInputSampleRates = C.kAudioCodecPropertyAvailableInputSampleRates;
    AudioCodecPropertyAvailableOutputSampleRates = C.kAudioCodecPropertyAvailableOutputSampleRates;
    AudioCodecPropertyAvailableBitRateRange = C.kAudioCodecPropertyAvailableBitRateRange;
    AudioCodecPropertyMinimumNumberInputPackets = C.kAudioCodecPropertyMinimumNumberInputPackets;
    AudioCodecPropertyMinimumNumberOutputPackets = C.kAudioCodecPropertyMinimumNumberOutputPackets;
    AudioCodecPropertyAvailableNumberChannels = C.kAudioCodecPropertyAvailableNumberChannels;
    AudioCodecPropertyDoesSampleRateConversion = C.kAudioCodecPropertyDoesSampleRateConversion;
    AudioCodecPropertyAvailableInputChannelLayoutTags = C.kAudioCodecPropertyAvailableInputChannelLayoutTags;
    AudioCodecPropertyAvailableOutputChannelLayoutTags = C.kAudioCodecPropertyAvailableOutputChannelLayoutTags;
    AudioCodecPropertyInputFormatsForOutputFormat = C.kAudioCodecPropertyInputFormatsForOutputFormat;
    AudioCodecPropertyOutputFormatsForInputFormat = C.kAudioCodecPropertyOutputFormatsForInputFormat;
    AudioCodecPropertyFormatInfo = C.kAudioCodecPropertyFormatInfo;
)
const (
    AudioCodecPropertyInputBufferSize = C.kAudioCodecPropertyInputBufferSize;
    AudioCodecPropertyPacketFrameSize = C.kAudioCodecPropertyPacketFrameSize;
    AudioCodecPropertyHasVariablePacketByteSizes = C.kAudioCodecPropertyHasVariablePacketByteSizes;
    AudioCodecPropertyMaximumPacketByteSize = C.kAudioCodecPropertyMaximumPacketByteSize;
    AudioCodecPropertyCurrentInputFormat = C.kAudioCodecPropertyCurrentInputFormat;
    AudioCodecPropertyCurrentOutputFormat = C.kAudioCodecPropertyCurrentOutputFormat;
    AudioCodecPropertyMagicCookie = C.kAudioCodecPropertyMagicCookie;
    AudioCodecPropertyUsedInputBufferSize = C.kAudioCodecPropertyUsedInputBufferSize;
    AudioCodecPropertyIsInitialized = C.kAudioCodecPropertyIsInitialized;
    AudioCodecPropertyCurrentTargetBitRate = C.kAudioCodecPropertyCurrentTargetBitRate;
    AudioCodecPropertyCurrentInputSampleRate = C.kAudioCodecPropertyCurrentInputSampleRate;
    AudioCodecPropertyCurrentOutputSampleRate = C.kAudioCodecPropertyCurrentOutputSampleRate;
    AudioCodecPropertyQualitySetting = C.kAudioCodecPropertyQualitySetting;
    AudioCodecPropertyApplicableBitRateRange = C.kAudioCodecPropertyApplicableBitRateRange;
    AudioCodecPropertyApplicableInputSampleRates = C.kAudioCodecPropertyApplicableInputSampleRates;
    AudioCodecPropertyApplicableOutputSampleRates = C.kAudioCodecPropertyApplicableOutputSampleRates;
    AudioCodecPropertyPaddedZeros = C.kAudioCodecPropertyPaddedZeros;
    AudioCodecPropertyPrimeMethod = C.kAudioCodecPropertyPrimeMethod;
    AudioCodecPropertyPrimeInfo = C.kAudioCodecPropertyPrimeInfo;
    AudioCodecPropertyCurrentInputChannelLayout = C.kAudioCodecPropertyCurrentInputChannelLayout;
    AudioCodecPropertyCurrentOutputChannelLayout = C.kAudioCodecPropertyCurrentOutputChannelLayout;
    AudioCodecPropertySettings = C.kAudioCodecPropertySettings;
    AudioCodecPropertyFormatList = C.kAudioCodecPropertyFormatList;
    AudioCodecPropertyBitRateControlMode = C.kAudioCodecPropertyBitRateControlMode;
    AudioCodecPropertySoundQualityForVBR = C.kAudioCodecPropertySoundQualityForVBR;
    AudioCodecPropertyDelayMode = C.kAudioCodecPropertyDelayMode;
)
const (
    AudioCodecQuality_Max = C.kAudioCodecQuality_Max;
    AudioCodecQuality_High = C.kAudioCodecQuality_High;
    AudioCodecQuality_Medium = C.kAudioCodecQuality_Medium;
    AudioCodecQuality_Low = C.kAudioCodecQuality_Low;
    AudioCodecQuality_Min = C.kAudioCodecQuality_Min;
)
const (
    AudioCodecPrimeMethod_Pre = C.kAudioCodecPrimeMethod_Pre;
    AudioCodecPrimeMethod_Normal = C.kAudioCodecPrimeMethod_Normal;
    AudioCodecPrimeMethod_None = C.kAudioCodecPrimeMethod_None;
)
const (
    AudioCodecBitRateControlMode_Constant = C.kAudioCodecBitRateControlMode_Constant;
    AudioCodecBitRateControlMode_LongTermAverage = C.kAudioCodecBitRateControlMode_LongTermAverage;
    AudioCodecBitRateControlMode_VariableConstrained = C.kAudioCodecBitRateControlMode_VariableConstrained;
    AudioCodecBitRateControlMode_Variable = C.kAudioCodecBitRateControlMode_Variable;
)
const (
    AudioCodecDelayMode_Compatibility = C.kAudioCodecDelayMode_Compatibility;
    AudioCodecDelayMode_Minimum = C.kAudioCodecDelayMode_Minimum;
    AudioCodecDelayMode_Optimal = C.kAudioCodecDelayMode_Optimal;
)
const (
    AudioSettingsFlags_ExpertParameter = C.kAudioSettingsFlags_ExpertParameter;
    AudioSettingsFlags_InvisibleParameter = C.kAudioSettingsFlags_InvisibleParameter;
    AudioSettingsFlags_MetaParameter = C.kAudioSettingsFlags_MetaParameter;
    AudioSettingsFlags_UserInterfaceParameter = C.kAudioSettingsFlags_UserInterfaceParameter;
)
const (
    AudioCodecProduceOutputPacketFailure = C.kAudioCodecProduceOutputPacketFailure;
    AudioCodecProduceOutputPacketSuccess = C.kAudioCodecProduceOutputPacketSuccess;
    AudioCodecProduceOutputPacketSuccessHasMore = C.kAudioCodecProduceOutputPacketSuccessHasMore;
    AudioCodecProduceOutputPacketNeedsMoreInputData = C.kAudioCodecProduceOutputPacketNeedsMoreInputData;
    AudioCodecProduceOutputPacketAtEOF = C.kAudioCodecProduceOutputPacketAtEOF;
)
const (
    AudioCodecGetPropertyInfoSelect = C.kAudioCodecGetPropertyInfoSelect;
    AudioCodecGetPropertySelect = C.kAudioCodecGetPropertySelect;
    AudioCodecSetPropertySelect = C.kAudioCodecSetPropertySelect;
    AudioCodecInitializeSelect = C.kAudioCodecInitializeSelect;
    AudioCodecUninitializeSelect = C.kAudioCodecUninitializeSelect;
    AudioCodecAppendInputDataSelect = C.kAudioCodecAppendInputDataSelect;
    AudioCodecProduceOutputDataSelect = C.kAudioCodecProduceOutputDataSelect;
    AudioCodecResetSelect = C.kAudioCodecResetSelect;
    AudioCodecAppendInputBufferListSelect = C.kAudioCodecAppendInputBufferListSelect;
    AudioCodecProduceOutputBufferListSelect = C.kAudioCodecProduceOutputBufferListSelect;
);
const (
    AudioCodecNoError = C.kAudioCodecNoError;
    AudioCodecUnspecifiedError = C.kAudioCodecUnspecifiedError;
    AudioCodecUnknownPropertyError = C.kAudioCodecUnknownPropertyError;
    AudioCodecBadPropertySizeError = C.kAudioCodecBadPropertySizeError;
    AudioCodecIllegalOperationError = C.kAudioCodecIllegalOperationError;
    AudioCodecUnsupportedFormatError = C.kAudioCodecUnsupportedFormatError;
    AudioCodecStateError = C.kAudioCodecStateError;
    AudioCodecNotEnoughBufferSpaceError = C.kAudioCodecNotEnoughBufferSpaceError;
)
const (
    AUPresetVersionKey = C.kAUPresetVersionKey
    AUPresetTypeKey = C.kAUPresetTypeKey
    AUPresetSubtypeKey = C.kAUPresetSubtypeKey
    AUPresetManufacturerKey = C.kAUPresetManufacturerKey
    AUPresetDataKey = C.kAUPresetDataKey
    AUPresetNameKey = C.kAUPresetNameKey
    AUPresetRenderQualityKey = C.kAUPresetRenderQualityKey
    AUPresetCPULoadKey = C.kAUPresetCPULoadKey
    AUPresetElementNameKey = C.kAUPresetElementNameKey
    AUPresetExternalFileRefs = C.kAUPresetExternalFileRefs
    AUPresetVSTDataKey = C.kAUPresetVSTDataKey
    AUPresetVSTPresetKey = C.kAUPresetVSTPresetKey
    AUPresetMASDataKey = C.kAUPresetMASDataKey
    AUPresetPartKey = C.kAUPresetPartKey
)

type AudioComponentDescription struct {
    ComponentType           int
    ComponentSubType        int
    ComponentManufacturer   int
    ComponentFlags          uint32
    ComponentFlagsMask      uint32
}

const (
    Audio_UnimplementedError = C.kAudio_UnimplementedError;
    Audio_FileNotFoundError = C.kAudio_FileNotFoundError;
    Audio_ParamError = C.kAudio_ParamError;
    Audio_MemFullError = C.kAudio_MemFullError;
)

type AudioValueRange struct {
    Minimum float64
    Maximum float64
}

type AudioValueTranslation struct {
    InputData  interface {}
    OutputData interface {}
}

type AudioBuffer struct {
    NumberChannels          uint32
    Data                    interface {}
}

type AudioBufferList []*AudioBuffer

// const (
//     AudioUnitSampleFractionBits = C.kAudioUnitSampleFractionBits;
// )

type AudioStreamBasicDescription struct {
    SampleRate              float64
    FormatID                uint32
    FormatFlags             uint32
    BytesPerPacket          uint32
    FramesPerPacket         uint32
    BytesPerFrame           uint32
    ChannelsPerFrame        uint32
    BitsPerChannel          uint32
}

func build_AudioStreamBasicDescription(desc AudioStreamBasicDescription) *C.AudioStreamBasicDescription {
    return &C.AudioStreamBasicDescription {
        C.Float64(desc.SampleRate),
        C.UInt32(desc.FormatID),
        C.UInt32(desc.FormatFlags),
        C.UInt32(desc.BytesPerPacket),
        C.UInt32(desc.FramesPerPacket),
        C.UInt32(desc.BytesPerFrame),
        C.UInt32(desc.ChannelsPerFrame),
        C.UInt32(desc.BitsPerChannel),
        0,
    }
}

const (
    AudioStreamAnyRate = C.kAudioStreamAnyRate;
)

const (
    AudioFormatLinearPCM = C.kAudioFormatLinearPCM;
    AudioFormatAC3 = C.kAudioFormatAC3;
    AudioFormat60958AC3 = C.kAudioFormat60958AC3;
    AudioFormatAppleIMA4 = C.kAudioFormatAppleIMA4;
    AudioFormatMPEG4AAC = C.kAudioFormatMPEG4AAC;
    AudioFormatMPEG4CELP = C.kAudioFormatMPEG4CELP;
    AudioFormatMPEG4HVXC = C.kAudioFormatMPEG4HVXC;
    AudioFormatMPEG4TwinVQ = C.kAudioFormatMPEG4TwinVQ;
    AudioFormatMACE3 = C.kAudioFormatMACE3;
    AudioFormatMACE6 = C.kAudioFormatMACE6;
    AudioFormatULaw = C.kAudioFormatULaw;
    AudioFormatALaw = C.kAudioFormatALaw;
    AudioFormatQDesign = C.kAudioFormatQDesign;
    AudioFormatQDesign2 = C.kAudioFormatQDesign2;
    AudioFormatQUALCOMM = C.kAudioFormatQUALCOMM;
    AudioFormatMPEGLayer1 = C.kAudioFormatMPEGLayer1;
    AudioFormatMPEGLayer2 = C.kAudioFormatMPEGLayer2;
    AudioFormatMPEGLayer3 = C.kAudioFormatMPEGLayer3;
    AudioFormatTimeCode = C.kAudioFormatTimeCode;
    AudioFormatMIDIStream = C.kAudioFormatMIDIStream;
    AudioFormatParameterValueStream = C.kAudioFormatParameterValueStream;
    AudioFormatAppleLossless = C.kAudioFormatAppleLossless;
    AudioFormatMPEG4AAC_HE = C.kAudioFormatMPEG4AAC_HE;
    AudioFormatMPEG4AAC_LD = C.kAudioFormatMPEG4AAC_LD;
    AudioFormatMPEG4AAC_ELD = C.kAudioFormatMPEG4AAC_ELD;
    AudioFormatMPEG4AAC_ELD_SBR = C.kAudioFormatMPEG4AAC_ELD_SBR;
    AudioFormatMPEG4AAC_ELD_V2 = C.kAudioFormatMPEG4AAC_ELD_V2;
    AudioFormatMPEG4AAC_HE_V2 = C.kAudioFormatMPEG4AAC_HE_V2;
    AudioFormatMPEG4AAC_Spatial = C.kAudioFormatMPEG4AAC_Spatial;
    AudioFormatAMR = C.kAudioFormatAMR;
    AudioFormatAudible = C.kAudioFormatAudible;
    AudioFormatiLBC = C.kAudioFormatiLBC;
    AudioFormatDVIIntelIMA = C.kAudioFormatDVIIntelIMA;
    AudioFormatMicrosoftGSM = C.kAudioFormatMicrosoftGSM;
    AudioFormatAES3 = C.kAudioFormatAES3;
)
const (
    AudioFormatFlagIsFloat = C.kAudioFormatFlagIsFloat;
    AudioFormatFlagIsBigEndian = C.kAudioFormatFlagIsBigEndian;
    AudioFormatFlagIsSignedInteger = C.kAudioFormatFlagIsSignedInteger;
    AudioFormatFlagIsPacked = C.kAudioFormatFlagIsPacked;
    AudioFormatFlagIsAlignedHigh = C.kAudioFormatFlagIsAlignedHigh;
    AudioFormatFlagIsNonInterleaved = C.kAudioFormatFlagIsNonInterleaved;
    AudioFormatFlagIsNonMixable = C.kAudioFormatFlagIsNonMixable;
    AudioFormatFlagsAreAllClear = C.kAudioFormatFlagsAreAllClear;
    LinearPCMFormatFlagIsFloat = C.kLinearPCMFormatFlagIsFloat;
    LinearPCMFormatFlagIsBigEndian = C.kLinearPCMFormatFlagIsBigEndian;
    LinearPCMFormatFlagIsSignedInteger = C.kLinearPCMFormatFlagIsSignedInteger;
    LinearPCMFormatFlagIsPacked = C.kLinearPCMFormatFlagIsPacked;
    LinearPCMFormatFlagIsAlignedHigh = C.kLinearPCMFormatFlagIsAlignedHigh;
    LinearPCMFormatFlagIsNonInterleaved = C.kLinearPCMFormatFlagIsNonInterleaved;
    LinearPCMFormatFlagIsNonMixable = C.kLinearPCMFormatFlagIsNonMixable;
    LinearPCMFormatFlagsSampleFractionShift = C.kLinearPCMFormatFlagsSampleFractionShift;
    LinearPCMFormatFlagsSampleFractionMask = C.kLinearPCMFormatFlagsSampleFractionMask;
    LinearPCMFormatFlagsAreAllClear = C.kLinearPCMFormatFlagsAreAllClear;
    AppleLosslessFormatFlag_16BitSourceData = C.kAppleLosslessFormatFlag_16BitSourceData;
    AppleLosslessFormatFlag_20BitSourceData = C.kAppleLosslessFormatFlag_20BitSourceData;
    AppleLosslessFormatFlag_24BitSourceData = C.kAppleLosslessFormatFlag_24BitSourceData;
    AppleLosslessFormatFlag_32BitSourceData = C.kAppleLosslessFormatFlag_32BitSourceData;
)

const (
    AudioFormatFlagsNativeEndian = C.kAudioFormatFlagsNativeEndian;
    AudioFormatFlagsCanonical = C.kAudioFormatFlagsCanonical;
    AudioFormatFlagsAudioUnitCanonical = C.kAudioFormatFlagsAudioUnitCanonical;
    AudioFormatFlagsNativeFloatPacked = C.kAudioFormatFlagsNativeFloatPacked;
)

type AudioStreamPacketDescription struct {
    StartOffset               int64
    VariableFramesInPacket    int32
    DataByteSize              int32
}

type SMPTETime struct {
    mSubframes                int16
    mSubframeDivisor          int16
    mCounter                  uint32
    mType                     uint32
    mFlags                    uint32
    mHours                    int16
    mMinutes                  int16
    mSeconds                  int16
    mFrames                   int16
}

const (
    SMPTETimeType24 = C.kSMPTETimeType24;
    SMPTETimeType25 = C.kSMPTETimeType25;
    SMPTETimeType30Drop = C.kSMPTETimeType30Drop;
    SMPTETimeType30 = C.kSMPTETimeType30;
    SMPTETimeType2997 = C.kSMPTETimeType2997;
    SMPTETimeType2997Drop = C.kSMPTETimeType2997Drop;
    SMPTETimeType60 = C.kSMPTETimeType60;
    SMPTETimeType5994 = C.kSMPTETimeType5994;
    SMPTETimeType60Drop = C.kSMPTETimeType60Drop;
    SMPTETimeType5994Drop = C.kSMPTETimeType5994Drop;
    SMPTETimeType50 = C.kSMPTETimeType50;
    SMPTETimeType2398 = C.kSMPTETimeType2398;
)

const (
    SMPTETimeValid = C.kSMPTETimeValid;
    SMPTETimeRunning = C.kSMPTETimeRunning;
)

type AudioTimeStamp struct {
    SampleTime              float64
    HostTime                uint64
    RateScalar              float64
    WordClockTime           uint64
    SMPTETime               SMPTETime
    Flags                   uint32
    Reserved                uint32
}

const (
    AudioTimeStampSampleTimeValid = C.kAudioTimeStampSampleTimeValid;
    AudioTimeStampHostTimeValid = C.kAudioTimeStampHostTimeValid;
    AudioTimeStampRateScalarValid = C.kAudioTimeStampRateScalarValid;
    AudioTimeStampWordClockTimeValid = C.kAudioTimeStampWordClockTimeValid;
    AudioTimeStampSMPTETimeValid = C.kAudioTimeStampSMPTETimeValid;
)

const (
    AudioTimeStampSampleHostTimeValid = C.kAudioTimeStampSampleHostTimeValid;
)

type AudioClassDescription struct {
    Type                uint32
    SubType             uint32
    Manufacturer        uint32
}

type AudioChannelDescription struct {
    ChannelLabel        uint32
    ChannelFlags        uint32
    Coordinates         [3]float32
}

type AudioChannelLayout struct {
    ChannelLayoutTag            uint32
    ChannelBitmap               uint32
    ChannelDescriptions         []AudioChannelDescription
}

const (
    AudioChannelLabel_Unknown = C.kAudioChannelLabel_Unknown;
    AudioChannelLabel_Unused = C.kAudioChannelLabel_Unused;
    AudioChannelLabel_UseCoordinates = C.kAudioChannelLabel_UseCoordinates;

    AudioChannelLabel_Left = C.kAudioChannelLabel_Left;
    AudioChannelLabel_Right = C.kAudioChannelLabel_Right;
    AudioChannelLabel_Center = C.kAudioChannelLabel_Center;
    AudioChannelLabel_LFEScreen = C.kAudioChannelLabel_LFEScreen;
    AudioChannelLabel_LeftSurround = C.kAudioChannelLabel_LeftSurround;
    AudioChannelLabel_RightSurround = C.kAudioChannelLabel_RightSurround;
    AudioChannelLabel_LeftCenter = C.kAudioChannelLabel_LeftCenter;
    AudioChannelLabel_RightCenter = C.kAudioChannelLabel_RightCenter;
    AudioChannelLabel_CenterSurround = C.kAudioChannelLabel_CenterSurround;
    AudioChannelLabel_LeftSurroundDirect = C.kAudioChannelLabel_LeftSurroundDirect;
    AudioChannelLabel_RightSurroundDirect = C.kAudioChannelLabel_RightSurroundDirect;
    AudioChannelLabel_TopCenterSurround = C.kAudioChannelLabel_TopCenterSurround;
    AudioChannelLabel_VerticalHeightLeft = C.kAudioChannelLabel_VerticalHeightLeft;
    AudioChannelLabel_VerticalHeightCenter = C.kAudioChannelLabel_VerticalHeightCenter;
    AudioChannelLabel_VerticalHeightRight = C.kAudioChannelLabel_VerticalHeightRight;

    AudioChannelLabel_TopBackLeft = C.kAudioChannelLabel_TopBackLeft;
    AudioChannelLabel_TopBackCenter = C.kAudioChannelLabel_TopBackCenter;
    AudioChannelLabel_TopBackRight = C.kAudioChannelLabel_TopBackRight;

    AudioChannelLabel_RearSurroundLeft = C.kAudioChannelLabel_RearSurroundLeft;
    AudioChannelLabel_RearSurroundRight = C.kAudioChannelLabel_RearSurroundRight;
    AudioChannelLabel_LeftWide = C.kAudioChannelLabel_LeftWide;
    AudioChannelLabel_RightWide = C.kAudioChannelLabel_RightWide;
    AudioChannelLabel_LFE2 = C.kAudioChannelLabel_LFE2;
    AudioChannelLabel_LeftTotal = C.kAudioChannelLabel_LeftTotal;
    AudioChannelLabel_RightTotal = C.kAudioChannelLabel_RightTotal;
    AudioChannelLabel_HearingImpaired = C.kAudioChannelLabel_HearingImpaired;
    AudioChannelLabel_Narration = C.kAudioChannelLabel_Narration;
    AudioChannelLabel_Mono = C.kAudioChannelLabel_Mono;
    AudioChannelLabel_DialogCentricMix = C.kAudioChannelLabel_DialogCentricMix;

    AudioChannelLabel_CenterSurroundDirect = C.kAudioChannelLabel_CenterSurroundDirect;
    AudioChannelLabel_Haptic = C.kAudioChannelLabel_Haptic;
    AudioChannelLabel_Ambisonic_W = C.kAudioChannelLabel_Ambisonic_W;
    AudioChannelLabel_Ambisonic_X = C.kAudioChannelLabel_Ambisonic_X;
    AudioChannelLabel_Ambisonic_Y = C.kAudioChannelLabel_Ambisonic_Y;
    AudioChannelLabel_Ambisonic_Z = C.kAudioChannelLabel_Ambisonic_Z;
    AudioChannelLabel_MS_Mid = C.kAudioChannelLabel_MS_Mid;
    AudioChannelLabel_MS_Side = C.kAudioChannelLabel_MS_Side;
    AudioChannelLabel_XY_X = C.kAudioChannelLabel_XY_X;
    AudioChannelLabel_XY_Y = C.kAudioChannelLabel_XY_Y;
    AudioChannelLabel_HeadphonesLeft = C.kAudioChannelLabel_HeadphonesLeft;
    AudioChannelLabel_HeadphonesRight = C.kAudioChannelLabel_HeadphonesRight;
    AudioChannelLabel_ClickTrack = C.kAudioChannelLabel_ClickTrack;
    AudioChannelLabel_ForeignLanguage = C.kAudioChannelLabel_ForeignLanguage;
    AudioChannelLabel_Discrete = C.kAudioChannelLabel_Discrete;
    AudioChannelLabel_Discrete_0 = C.kAudioChannelLabel_Discrete_0;
    AudioChannelLabel_Discrete_1 = C.kAudioChannelLabel_Discrete_1;
    AudioChannelLabel_Discrete_2 = C.kAudioChannelLabel_Discrete_2;
    AudioChannelLabel_Discrete_3 = C.kAudioChannelLabel_Discrete_3;
    AudioChannelLabel_Discrete_4 = C.kAudioChannelLabel_Discrete_4;
    AudioChannelLabel_Discrete_5 = C.kAudioChannelLabel_Discrete_5;
    AudioChannelLabel_Discrete_6 = C.kAudioChannelLabel_Discrete_6;
    AudioChannelLabel_Discrete_7 = C.kAudioChannelLabel_Discrete_7;
    AudioChannelLabel_Discrete_8 = C.kAudioChannelLabel_Discrete_8;
    AudioChannelLabel_Discrete_9 = C.kAudioChannelLabel_Discrete_9;
    AudioChannelLabel_Discrete_10 = C.kAudioChannelLabel_Discrete_10;
    AudioChannelLabel_Discrete_11 = C.kAudioChannelLabel_Discrete_11;
    AudioChannelLabel_Discrete_12 = C.kAudioChannelLabel_Discrete_12;
    AudioChannelLabel_Discrete_13 = C.kAudioChannelLabel_Discrete_13;
    AudioChannelLabel_Discrete_14 = C.kAudioChannelLabel_Discrete_14;
    AudioChannelLabel_Discrete_15 = C.kAudioChannelLabel_Discrete_15;
    AudioChannelLabel_Discrete_65535 = C.kAudioChannelLabel_Discrete_65535;
)

const (
    AudioChannelBit_Left = C.kAudioChannelBit_Left;
    AudioChannelBit_Right = C.kAudioChannelBit_Right;
    AudioChannelBit_Center = C.kAudioChannelBit_Center;
    AudioChannelBit_LFEScreen = C.kAudioChannelBit_LFEScreen;
    AudioChannelBit_LeftSurround = C.kAudioChannelBit_LeftSurround;
    AudioChannelBit_RightSurround = C.kAudioChannelBit_RightSurround;
    AudioChannelBit_LeftCenter = C.kAudioChannelBit_LeftCenter;
    AudioChannelBit_RightCenter = C.kAudioChannelBit_RightCenter;
    AudioChannelBit_CenterSurround = C.kAudioChannelBit_CenterSurround;
    AudioChannelBit_LeftSurroundDirect = C.kAudioChannelBit_LeftSurroundDirect;
    AudioChannelBit_RightSurroundDirect = C.kAudioChannelBit_RightSurroundDirect;
    AudioChannelBit_TopCenterSurround = C.kAudioChannelBit_TopCenterSurround;
    AudioChannelBit_VerticalHeightLeft = C.kAudioChannelBit_VerticalHeightLeft;
    AudioChannelBit_VerticalHeightCenter = C.kAudioChannelBit_VerticalHeightCenter;
    AudioChannelBit_VerticalHeightRight = C.kAudioChannelBit_VerticalHeightRight;
    AudioChannelBit_TopBackLeft = C.kAudioChannelBit_TopBackLeft;
    AudioChannelBit_TopBackCenter = C.kAudioChannelBit_TopBackCenter;
    AudioChannelBit_TopBackRight = C.kAudioChannelBit_TopBackRight;
)

const (
    AudioChannelFlags_AllOff = C.kAudioChannelFlags_AllOff;
    AudioChannelFlags_RectangularCoordinates = C.kAudioChannelFlags_RectangularCoordinates;
    AudioChannelFlags_SphericalCoordinates = C.kAudioChannelFlags_SphericalCoordinates;
    AudioChannelFlags_Meters = C.kAudioChannelFlags_Meters;
)

const (
    AudioChannelCoordinates_LeftRight = C.kAudioChannelCoordinates_LeftRight;
    AudioChannelCoordinates_BackFront = C.kAudioChannelCoordinates_BackFront;
    AudioChannelCoordinates_DownUp = C.kAudioChannelCoordinates_DownUp;
    AudioChannelCoordinates_Azimuth = C.kAudioChannelCoordinates_Azimuth;
    AudioChannelCoordinates_Elevation = C.kAudioChannelCoordinates_Elevation;
    AudioChannelCoordinates_Distance = C.kAudioChannelCoordinates_Distance;
)

const (
    AudioChannelLayoutTag_UseChannelDescriptions = C.kAudioChannelLayoutTag_UseChannelDescriptions;
    AudioChannelLayoutTag_UseChannelBitmap = C.kAudioChannelLayoutTag_UseChannelBitmap;

    AudioChannelLayoutTag_Mono = C.kAudioChannelLayoutTag_Mono;
    AudioChannelLayoutTag_Stereo = C.kAudioChannelLayoutTag_Stereo;
    AudioChannelLayoutTag_StereoHeadphones = C.kAudioChannelLayoutTag_StereoHeadphones;
    AudioChannelLayoutTag_MatrixStereo = C.kAudioChannelLayoutTag_MatrixStereo;
    AudioChannelLayoutTag_MidSide = C.kAudioChannelLayoutTag_MidSide;
    AudioChannelLayoutTag_XY = C.kAudioChannelLayoutTag_XY;
    AudioChannelLayoutTag_Binaural = C.kAudioChannelLayoutTag_Binaural;
    AudioChannelLayoutTag_Ambisonic_B_Format = C.kAudioChannelLayoutTag_Ambisonic_B_Format;

    AudioChannelLayoutTag_Quadraphonic = C.kAudioChannelLayoutTag_Quadraphonic;
    AudioChannelLayoutTag_Pentagonal = C.kAudioChannelLayoutTag_Pentagonal;
    AudioChannelLayoutTag_Hexagonal = C.kAudioChannelLayoutTag_Hexagonal;
    AudioChannelLayoutTag_Octagonal = C.kAudioChannelLayoutTag_Octagonal;
    AudioChannelLayoutTag_Cube = C.kAudioChannelLayoutTag_Cube;
    AudioChannelLayoutTag_MPEG_1_0 = C.kAudioChannelLayoutTag_MPEG_1_0;
    AudioChannelLayoutTag_MPEG_2_0 = C.kAudioChannelLayoutTag_MPEG_2_0;
    AudioChannelLayoutTag_MPEG_3_0_A = C.kAudioChannelLayoutTag_MPEG_3_0_A;
    AudioChannelLayoutTag_MPEG_3_0_B = C.kAudioChannelLayoutTag_MPEG_3_0_B;
    AudioChannelLayoutTag_MPEG_4_0_A = C.kAudioChannelLayoutTag_MPEG_4_0_A;
    AudioChannelLayoutTag_MPEG_4_0_B = C.kAudioChannelLayoutTag_MPEG_4_0_B;
    AudioChannelLayoutTag_MPEG_5_0_A = C.kAudioChannelLayoutTag_MPEG_5_0_A;
    AudioChannelLayoutTag_MPEG_5_0_B = C.kAudioChannelLayoutTag_MPEG_5_0_B;
    AudioChannelLayoutTag_MPEG_5_0_C = C.kAudioChannelLayoutTag_MPEG_5_0_C;
    AudioChannelLayoutTag_MPEG_5_0_D = C.kAudioChannelLayoutTag_MPEG_5_0_D;
    AudioChannelLayoutTag_MPEG_5_1_A = C.kAudioChannelLayoutTag_MPEG_5_1_A;
    AudioChannelLayoutTag_MPEG_5_1_B = C.kAudioChannelLayoutTag_MPEG_5_1_B;
    AudioChannelLayoutTag_MPEG_5_1_C = C.kAudioChannelLayoutTag_MPEG_5_1_C;
    AudioChannelLayoutTag_MPEG_5_1_D = C.kAudioChannelLayoutTag_MPEG_5_1_D;
    AudioChannelLayoutTag_MPEG_6_1_A = C.kAudioChannelLayoutTag_MPEG_6_1_A;
    AudioChannelLayoutTag_MPEG_7_1_A = C.kAudioChannelLayoutTag_MPEG_7_1_A;
    AudioChannelLayoutTag_MPEG_7_1_B = C.kAudioChannelLayoutTag_MPEG_7_1_B;
    AudioChannelLayoutTag_MPEG_7_1_C = C.kAudioChannelLayoutTag_MPEG_7_1_C;
    AudioChannelLayoutTag_Emagic_Default_7_1 = C.kAudioChannelLayoutTag_Emagic_Default_7_1;
    AudioChannelLayoutTag_SMPTE_DTV = C.kAudioChannelLayoutTag_SMPTE_DTV;
    AudioChannelLayoutTag_ITU_1_0 = C.kAudioChannelLayoutTag_ITU_1_0;
    AudioChannelLayoutTag_ITU_2_0 = C.kAudioChannelLayoutTag_ITU_2_0;

    AudioChannelLayoutTag_ITU_2_1 = C.kAudioChannelLayoutTag_ITU_2_1;
    AudioChannelLayoutTag_ITU_2_2 = C.kAudioChannelLayoutTag_ITU_2_2;
    AudioChannelLayoutTag_ITU_3_0 = C.kAudioChannelLayoutTag_ITU_3_0;
    AudioChannelLayoutTag_ITU_3_1 = C.kAudioChannelLayoutTag_ITU_3_1;

    AudioChannelLayoutTag_ITU_3_2 = C.kAudioChannelLayoutTag_ITU_3_2;
    AudioChannelLayoutTag_ITU_3_2_1 = C.kAudioChannelLayoutTag_ITU_3_2_1;
    AudioChannelLayoutTag_ITU_3_4_1 = C.kAudioChannelLayoutTag_ITU_3_4_1;

    AudioChannelLayoutTag_DVD_0 = C.kAudioChannelLayoutTag_DVD_0;
    AudioChannelLayoutTag_DVD_1 = C.kAudioChannelLayoutTag_DVD_1;
    AudioChannelLayoutTag_DVD_2 = C.kAudioChannelLayoutTag_DVD_2;
    AudioChannelLayoutTag_DVD_3 = C.kAudioChannelLayoutTag_DVD_3;
    AudioChannelLayoutTag_DVD_4 = C.kAudioChannelLayoutTag_DVD_4;
    AudioChannelLayoutTag_DVD_5 = C.kAudioChannelLayoutTag_DVD_5;
    AudioChannelLayoutTag_DVD_6 = C.kAudioChannelLayoutTag_DVD_6;
    AudioChannelLayoutTag_DVD_7 = C.kAudioChannelLayoutTag_DVD_7;
    AudioChannelLayoutTag_DVD_8 = C.kAudioChannelLayoutTag_DVD_8;
    AudioChannelLayoutTag_DVD_9 = C.kAudioChannelLayoutTag_DVD_9;
    AudioChannelLayoutTag_DVD_10 = C.kAudioChannelLayoutTag_DVD_10;
    AudioChannelLayoutTag_DVD_11 = C.kAudioChannelLayoutTag_DVD_11;
    AudioChannelLayoutTag_DVD_12 = C.kAudioChannelLayoutTag_DVD_12;
    AudioChannelLayoutTag_DVD_13 = C.kAudioChannelLayoutTag_DVD_13;
    AudioChannelLayoutTag_DVD_14 = C.kAudioChannelLayoutTag_DVD_14;
    AudioChannelLayoutTag_DVD_15 = C.kAudioChannelLayoutTag_DVD_15;
    AudioChannelLayoutTag_DVD_16 = C.kAudioChannelLayoutTag_DVD_16;
    AudioChannelLayoutTag_DVD_17 = C.kAudioChannelLayoutTag_DVD_17;
    AudioChannelLayoutTag_DVD_18 = C.kAudioChannelLayoutTag_DVD_18;
    AudioChannelLayoutTag_DVD_19 = C.kAudioChannelLayoutTag_DVD_19;
    AudioChannelLayoutTag_DVD_20 = C.kAudioChannelLayoutTag_DVD_20;

    AudioChannelLayoutTag_AudioUnit_4 = C.kAudioChannelLayoutTag_AudioUnit_4;
    AudioChannelLayoutTag_AudioUnit_5 = C.kAudioChannelLayoutTag_AudioUnit_5;
    AudioChannelLayoutTag_AudioUnit_6 = C.kAudioChannelLayoutTag_AudioUnit_6;
    AudioChannelLayoutTag_AudioUnit_8 = C.kAudioChannelLayoutTag_AudioUnit_8;
    AudioChannelLayoutTag_AudioUnit_5_0 = C.kAudioChannelLayoutTag_AudioUnit_5_0;
    AudioChannelLayoutTag_AudioUnit_6_0 = C.kAudioChannelLayoutTag_AudioUnit_6_0;
    AudioChannelLayoutTag_AudioUnit_7_0 = C.kAudioChannelLayoutTag_AudioUnit_7_0;
    AudioChannelLayoutTag_AudioUnit_7_0_Front = C.kAudioChannelLayoutTag_AudioUnit_7_0_Front;
    AudioChannelLayoutTag_AudioUnit_5_1 = C.kAudioChannelLayoutTag_AudioUnit_5_1;
    AudioChannelLayoutTag_AudioUnit_6_1 = C.kAudioChannelLayoutTag_AudioUnit_6_1;
    AudioChannelLayoutTag_AudioUnit_7_1 = C.kAudioChannelLayoutTag_AudioUnit_7_1;
    AudioChannelLayoutTag_AudioUnit_7_1_Front = C.kAudioChannelLayoutTag_AudioUnit_7_1_Front;

    AudioChannelLayoutTag_AAC_3_0 = C.kAudioChannelLayoutTag_AAC_3_0;
    AudioChannelLayoutTag_AAC_Quadraphonic = C.kAudioChannelLayoutTag_AAC_Quadraphonic;
    AudioChannelLayoutTag_AAC_4_0 = C.kAudioChannelLayoutTag_AAC_4_0;
    AudioChannelLayoutTag_AAC_5_0 = C.kAudioChannelLayoutTag_AAC_5_0;
    AudioChannelLayoutTag_AAC_5_1 = C.kAudioChannelLayoutTag_AAC_5_1;
    AudioChannelLayoutTag_AAC_6_0 = C.kAudioChannelLayoutTag_AAC_6_0;
    AudioChannelLayoutTag_AAC_6_1 = C.kAudioChannelLayoutTag_AAC_6_1;
    AudioChannelLayoutTag_AAC_7_0 = C.kAudioChannelLayoutTag_AAC_7_0;
    AudioChannelLayoutTag_AAC_7_1 = C.kAudioChannelLayoutTag_AAC_7_1;
    AudioChannelLayoutTag_AAC_7_1_B = C.kAudioChannelLayoutTag_AAC_7_1_B;
    AudioChannelLayoutTag_AAC_Octagonal = C.kAudioChannelLayoutTag_AAC_Octagonal;

    AudioChannelLayoutTag_TMH_10_2_std = C.kAudioChannelLayoutTag_TMH_10_2_std;
    AudioChannelLayoutTag_TMH_10_2_full = C.kAudioChannelLayoutTag_TMH_10_2_full;

    AudioChannelLayoutTag_AC3_1_0_1 = C.kAudioChannelLayoutTag_AC3_1_0_1;
    AudioChannelLayoutTag_AC3_3_0 = C.kAudioChannelLayoutTag_AC3_3_0;
    AudioChannelLayoutTag_AC3_3_1 = C.kAudioChannelLayoutTag_AC3_3_1;
    AudioChannelLayoutTag_AC3_3_0_1 = C.kAudioChannelLayoutTag_AC3_3_0_1;
    AudioChannelLayoutTag_AC3_2_1_1 = C.kAudioChannelLayoutTag_AC3_2_1_1;
    AudioChannelLayoutTag_AC3_3_1_1 = C.kAudioChannelLayoutTag_AC3_3_1_1;

    AudioChannelLayoutTag_EAC_6_0_A = C.kAudioChannelLayoutTag_EAC_6_0_A;
    AudioChannelLayoutTag_EAC_7_0_A = C.kAudioChannelLayoutTag_EAC_7_0_A;

    AudioChannelLayoutTag_EAC3_6_1_A = C.kAudioChannelLayoutTag_EAC3_6_1_A;
    AudioChannelLayoutTag_EAC3_6_1_B = C.kAudioChannelLayoutTag_EAC3_6_1_B;
    AudioChannelLayoutTag_EAC3_6_1_C = C.kAudioChannelLayoutTag_EAC3_6_1_C;
    AudioChannelLayoutTag_EAC3_7_1_A = C.kAudioChannelLayoutTag_EAC3_7_1_A;
    AudioChannelLayoutTag_EAC3_7_1_B = C.kAudioChannelLayoutTag_EAC3_7_1_B;
    AudioChannelLayoutTag_EAC3_7_1_C = C.kAudioChannelLayoutTag_EAC3_7_1_C;
    AudioChannelLayoutTag_EAC3_7_1_D = C.kAudioChannelLayoutTag_EAC3_7_1_D;
    AudioChannelLayoutTag_EAC3_7_1_E = C.kAudioChannelLayoutTag_EAC3_7_1_E;

    AudioChannelLayoutTag_EAC3_7_1_F = C.kAudioChannelLayoutTag_EAC3_7_1_F;
    AudioChannelLayoutTag_EAC3_7_1_G = C.kAudioChannelLayoutTag_EAC3_7_1_G;
    AudioChannelLayoutTag_EAC3_7_1_H = C.kAudioChannelLayoutTag_EAC3_7_1_H;

    AudioChannelLayoutTag_DTS_3_1 = C.kAudioChannelLayoutTag_DTS_3_1;
    AudioChannelLayoutTag_DTS_4_1 = C.kAudioChannelLayoutTag_DTS_4_1;
    AudioChannelLayoutTag_DTS_6_0_A = C.kAudioChannelLayoutTag_DTS_6_0_A;
    AudioChannelLayoutTag_DTS_6_0_B = C.kAudioChannelLayoutTag_DTS_6_0_B;
    AudioChannelLayoutTag_DTS_6_0_C = C.kAudioChannelLayoutTag_DTS_6_0_C;
    AudioChannelLayoutTag_DTS_6_1_A = C.kAudioChannelLayoutTag_DTS_6_1_A;
    AudioChannelLayoutTag_DTS_6_1_B = C.kAudioChannelLayoutTag_DTS_6_1_B;
    AudioChannelLayoutTag_DTS_6_1_C = C.kAudioChannelLayoutTag_DTS_6_1_C;
    AudioChannelLayoutTag_DTS_7_0 = C.kAudioChannelLayoutTag_DTS_7_0;
    AudioChannelLayoutTag_DTS_7_1 = C.kAudioChannelLayoutTag_DTS_7_1;
    AudioChannelLayoutTag_DTS_8_0_A = C.kAudioChannelLayoutTag_DTS_8_0_A;
    AudioChannelLayoutTag_DTS_8_0_B = C.kAudioChannelLayoutTag_DTS_8_0_B;
    AudioChannelLayoutTag_DTS_8_1_A = C.kAudioChannelLayoutTag_DTS_8_1_A;
    AudioChannelLayoutTag_DTS_8_1_B = C.kAudioChannelLayoutTag_DTS_8_1_B;
    AudioChannelLayoutTag_DTS_6_1_D = C.kAudioChannelLayoutTag_DTS_6_1_D;
    AudioChannelLayoutTag_DiscreteInOrder = C.kAudioChannelLayoutTag_DiscreteInOrder;
    AudioChannelLayoutTag_Unknown = C.kAudioChannelLayoutTag_Unknown;
)

const (
    MPEG4Object_AAC_Main = C.kMPEG4Object_AAC_Main;
    MPEG4Object_AAC_LC = C.kMPEG4Object_AAC_LC;
    MPEG4Object_AAC_SSR = C.kMPEG4Object_AAC_SSR;
    MPEG4Object_AAC_LTP = C.kMPEG4Object_AAC_LTP;
    MPEG4Object_AAC_SBR = C.kMPEG4Object_AAC_SBR;
    MPEG4Object_AAC_Scalable = C.kMPEG4Object_AAC_Scalable;
    MPEG4Object_TwinVQ = C.kMPEG4Object_TwinVQ;
    MPEG4Object_CELP = C.kMPEG4Object_CELP;
    MPEG4Object_HVXC = C.kMPEG4Object_HVXC;
)

const (
    AudioUnitScope_Global = C.kAudioUnitScope_Global;
    AudioUnitScope_Input = C.kAudioUnitScope_Input;
    AudioUnitScope_Output = C.kAudioUnitScope_Output;
    AudioUnitScope_Group = C.kAudioUnitScope_Group;
    AudioUnitScope_Part = C.kAudioUnitScope_Part;
    AudioUnitScope_Note = C.kAudioUnitScope_Note;
    AudioUnitScope_Layer = C.kAudioUnitScope_Layer;
    AudioUnitScope_LayerItem = C.kAudioUnitScope_LayerItem;
)

const (
    AudioUnitProperty_ClassInfo = C.kAudioUnitProperty_ClassInfo;
    AudioUnitProperty_MakeConnection = C.kAudioUnitProperty_MakeConnection;
    AudioUnitProperty_SampleRate = C.kAudioUnitProperty_SampleRate;
    AudioUnitProperty_ParameterList = C.kAudioUnitProperty_ParameterList;
    AudioUnitProperty_ParameterInfo = C.kAudioUnitProperty_ParameterInfo;
    AudioUnitProperty_CPULoad = C.kAudioUnitProperty_CPULoad;
    AudioUnitProperty_StreamFormat = C.kAudioUnitProperty_StreamFormat;
    AudioUnitProperty_ElementCount = C.kAudioUnitProperty_ElementCount;
    AudioUnitProperty_Latency = C.kAudioUnitProperty_Latency;
    AudioUnitProperty_SupportedNumChannels = C.kAudioUnitProperty_SupportedNumChannels;
    AudioUnitProperty_MaximumFramesPerSlice = C.kAudioUnitProperty_MaximumFramesPerSlice;
    AudioUnitProperty_ParameterValueStrings = C.kAudioUnitProperty_ParameterValueStrings;
    AudioUnitProperty_AudioChannelLayout = C.kAudioUnitProperty_AudioChannelLayout;
    AudioUnitProperty_TailTime = C.kAudioUnitProperty_TailTime;
    AudioUnitProperty_BypassEffect = C.kAudioUnitProperty_BypassEffect;
    AudioUnitProperty_LastRenderError = C.kAudioUnitProperty_LastRenderError;
    AudioUnitProperty_SetRenderCallback = C.kAudioUnitProperty_SetRenderCallback;
    AudioUnitProperty_FactoryPresets = C.kAudioUnitProperty_FactoryPresets;
    AudioUnitProperty_RenderQuality = C.kAudioUnitProperty_RenderQuality;
    AudioUnitProperty_InPlaceProcessing = C.kAudioUnitProperty_InPlaceProcessing;
    AudioUnitProperty_ElementName = C.kAudioUnitProperty_ElementName;
    AudioUnitProperty_SupportedChannelLayoutTags = C.kAudioUnitProperty_SupportedChannelLayoutTags;
    AudioUnitProperty_PresentPreset = C.kAudioUnitProperty_PresentPreset;
    AudioUnitProperty_ShouldAllocateBuffer = C.kAudioUnitProperty_ShouldAllocateBuffer;
    AudioUnitProperty_ParameterHistoryInfo = C.kAudioUnitProperty_ParameterHistoryInfo;
    AudioUnitProperty_FastDispatch = C.kAudioUnitProperty_FastDispatch;
    AudioUnitProperty_SetExternalBuffer = C.kAudioUnitProperty_SetExternalBuffer;
    AudioUnitProperty_GetUIComponentList = C.kAudioUnitProperty_GetUIComponentList;
    AudioUnitProperty_ContextName = C.kAudioUnitProperty_ContextName;
    AudioUnitProperty_HostCallbacks = C.kAudioUnitProperty_HostCallbacks;
    AudioUnitProperty_CocoaUI = C.kAudioUnitProperty_CocoaUI;
    AudioUnitProperty_ParameterIDName = C.kAudioUnitProperty_ParameterIDName;
    AudioUnitProperty_ParameterClumpName = C.kAudioUnitProperty_ParameterClumpName;
    AudioUnitProperty_ParameterStringFromValue = C.kAudioUnitProperty_ParameterStringFromValue;
    AudioUnitProperty_OfflineRender = C.kAudioUnitProperty_OfflineRender;
    AudioUnitProperty_ParameterValueFromString = C.kAudioUnitProperty_ParameterValueFromString;
    AudioUnitProperty_IconLocation = C.kAudioUnitProperty_IconLocation;
    AudioUnitProperty_PresentationLatency = C.kAudioUnitProperty_PresentationLatency;
    AudioUnitProperty_DependentParameters = C.kAudioUnitProperty_DependentParameters;
    AudioUnitProperty_AUHostIdentifier = C.kAudioUnitProperty_AUHostIdentifier;
    AudioUnitProperty_MIDIOutputCallbackInfo = C.kAudioUnitProperty_MIDIOutputCallbackInfo;
    AudioUnitProperty_MIDIOutputCallback = C.kAudioUnitProperty_MIDIOutputCallback;
    AudioUnitProperty_InputSamplesInOutput = C.kAudioUnitProperty_InputSamplesInOutput;
    AudioUnitProperty_ClassInfoFromDocument = C.kAudioUnitProperty_ClassInfoFromDocument;
    AudioUnitProperty_FrequencyResponse = C.kAudioUnitProperty_FrequencyResponse;
)

const (
    RenderQuality_Max = C.kRenderQuality_Max;
    RenderQuality_High = C.kRenderQuality_High;
    RenderQuality_Medium = C.kRenderQuality_Medium;
    RenderQuality_Low = C.kRenderQuality_Low;
    RenderQuality_Min = C.kRenderQuality_Min;
)

const (
    AudioUnitParameterUnit_Generic = C.kAudioUnitParameterUnit_Generic;
    AudioUnitParameterUnit_Indexed = C.kAudioUnitParameterUnit_Indexed;
    AudioUnitParameterUnit_Boolean = C.kAudioUnitParameterUnit_Boolean;
    AudioUnitParameterUnit_Percent = C.kAudioUnitParameterUnit_Percent;
    AudioUnitParameterUnit_Seconds = C.kAudioUnitParameterUnit_Seconds;
    AudioUnitParameterUnit_SampleFrames = C.kAudioUnitParameterUnit_SampleFrames;
    AudioUnitParameterUnit_Phase = C.kAudioUnitParameterUnit_Phase;
    AudioUnitParameterUnit_Rate = C.kAudioUnitParameterUnit_Rate;
    AudioUnitParameterUnit_Hertz = C.kAudioUnitParameterUnit_Hertz;
    AudioUnitParameterUnit_Cents = C.kAudioUnitParameterUnit_Cents;
    AudioUnitParameterUnit_RelativeSemiTones = C.kAudioUnitParameterUnit_RelativeSemiTones;
    AudioUnitParameterUnit_MIDINoteNumber = C.kAudioUnitParameterUnit_MIDINoteNumber;
    AudioUnitParameterUnit_MIDIController = C.kAudioUnitParameterUnit_MIDIController;
    AudioUnitParameterUnit_Decibels = C.kAudioUnitParameterUnit_Decibels;
    AudioUnitParameterUnit_LinearGain = C.kAudioUnitParameterUnit_LinearGain;
    AudioUnitParameterUnit_Degrees = C.kAudioUnitParameterUnit_Degrees;
    AudioUnitParameterUnit_EqualPowerCrossfade = C.kAudioUnitParameterUnit_EqualPowerCrossfade;
    AudioUnitParameterUnit_MixerFaderCurve1 = C.kAudioUnitParameterUnit_MixerFaderCurve1;
    AudioUnitParameterUnit_Pan = C.kAudioUnitParameterUnit_Pan;
    AudioUnitParameterUnit_Meters = C.kAudioUnitParameterUnit_Meters;
    AudioUnitParameterUnit_AbsoluteCents = C.kAudioUnitParameterUnit_AbsoluteCents;
    AudioUnitParameterUnit_Octaves = C.kAudioUnitParameterUnit_Octaves;
    AudioUnitParameterUnit_BPM = C.kAudioUnitParameterUnit_BPM;
    AudioUnitParameterUnit_Beats = C.kAudioUnitParameterUnit_Beats;
    AudioUnitParameterUnit_Milliseconds = C.kAudioUnitParameterUnit_Milliseconds;
    AudioUnitParameterUnit_Ratio = C.kAudioUnitParameterUnit_Ratio;
    AudioUnitParameterUnit_CustomUnit = C.kAudioUnitParameterUnit_CustomUnit;
)

const (
    AudioUnitParameterFlag_CFNameRelease = C.kAudioUnitParameterFlag_CFNameRelease;

    AudioUnitParameterFlag_PlotHistory = C.kAudioUnitParameterFlag_PlotHistory;
    AudioUnitParameterFlag_MeterReadOnly = C.kAudioUnitParameterFlag_MeterReadOnly;
    AudioUnitParameterFlag_DisplayMask = C.kAudioUnitParameterFlag_DisplayMask;
    AudioUnitParameterFlag_DisplaySquareRoot = C.kAudioUnitParameterFlag_DisplaySquareRoot;
    AudioUnitParameterFlag_DisplaySquared = C.kAudioUnitParameterFlag_DisplaySquared;
    AudioUnitParameterFlag_DisplayCubed = C.kAudioUnitParameterFlag_DisplayCubed;
    AudioUnitParameterFlag_DisplayCubeRoot = C.kAudioUnitParameterFlag_DisplayCubeRoot;
    AudioUnitParameterFlag_DisplayExponential = C.kAudioUnitParameterFlag_DisplayExponential;

    AudioUnitParameterFlag_HasClump = C.kAudioUnitParameterFlag_HasClump;
    AudioUnitParameterFlag_ValuesHaveStrings = C.kAudioUnitParameterFlag_ValuesHaveStrings;
    AudioUnitParameterFlag_DisplayLogarithmic = C.kAudioUnitParameterFlag_DisplayLogarithmic;
    AudioUnitParameterFlag_IsHighResolution = C.kAudioUnitParameterFlag_IsHighResolution;
    AudioUnitParameterFlag_NonRealTime = C.kAudioUnitParameterFlag_NonRealTime;
    AudioUnitParameterFlag_CanRamp = C.kAudioUnitParameterFlag_CanRamp;
    AudioUnitParameterFlag_ExpertMode = C.kAudioUnitParameterFlag_ExpertMode;
    AudioUnitParameterFlag_HasCFNameString = C.kAudioUnitParameterFlag_HasCFNameString;
    AudioUnitParameterFlag_IsGlobalMeta = C.kAudioUnitParameterFlag_IsGlobalMeta;
    AudioUnitParameterFlag_IsElementMeta = C.kAudioUnitParameterFlag_IsElementMeta;
    AudioUnitParameterFlag_IsReadable = C.kAudioUnitParameterFlag_IsReadable;
    AudioUnitParameterFlag_IsWritable = C.kAudioUnitParameterFlag_IsWritable;
)

const (
    AudioUnitClumpID_System = C.kAudioUnitClumpID_System;
)

const (
    AudioUnitParameterName_Full = C.kAudioUnitParameterName_Full;
)

const (
    AudioOutputUnitProperty_IsRunning = C.kAudioOutputUnitProperty_IsRunning;
)

const (
    AudioUnitProperty_AllParameterMIDIMappings = C.kAudioUnitProperty_AllParameterMIDIMappings;
    AudioUnitProperty_AddParameterMIDIMapping = C.kAudioUnitProperty_AddParameterMIDIMapping;
    AudioUnitProperty_RemoveParameterMIDIMapping = C.kAudioUnitProperty_RemoveParameterMIDIMapping;
    AudioUnitProperty_HotMapParameterMIDIMapping = C.kAudioUnitProperty_HotMapParameterMIDIMapping;
)

const (
    AUParameterMIDIMapping_AnyChannelFlag = C.kAUParameterMIDIMapping_AnyChannelFlag;
    AUParameterMIDIMapping_AnyNoteFlag = C.kAUParameterMIDIMapping_AnyNoteFlag;
    AUParameterMIDIMapping_SubRange = C.kAUParameterMIDIMapping_SubRange;
    AUParameterMIDIMapping_Toggle = C.kAUParameterMIDIMapping_Toggle;
    AUParameterMIDIMapping_Bipolar = C.kAUParameterMIDIMapping_Bipolar;
    AUParameterMIDIMapping_Bipolar_On = C.kAUParameterMIDIMapping_Bipolar_On;
)

const (
    MusicDeviceProperty_InstrumentCount = C.kMusicDeviceProperty_InstrumentCount;
    MusicDeviceProperty_MIDIXMLNames = C.kMusicDeviceProperty_MIDIXMLNames;
    MusicDeviceProperty_PartGroup = C.kMusicDeviceProperty_PartGroup;
    MusicDeviceProperty_DualSchedulingMode = C.kMusicDeviceProperty_DualSchedulingMode;
    MusicDeviceProperty_SupportsStartStopNote = C.kMusicDeviceProperty_SupportsStartStopNote;
)

const (
    MusicDeviceSampleFrameMask_SampleOffset = C.kMusicDeviceSampleFrameMask_SampleOffset;
    MusicDeviceSampleFrameMask_IsScheduled = C.kMusicDeviceSampleFrameMask_IsScheduled;
)

const (
    AudioUnitOfflineProperty_InputSize = C.kAudioUnitOfflineProperty_InputSize;
    AudioUnitOfflineProperty_OutputSize = C.kAudioUnitOfflineProperty_OutputSize;
    AudioUnitOfflineProperty_StartOffset = C.kAudioUnitOfflineProperty_StartOffset;
    AudioUnitOfflineProperty_PreflightRequirements = C.kAudioUnitOfflineProperty_PreflightRequirements;
    AudioUnitOfflineProperty_PreflightName = C.kAudioUnitOfflineProperty_PreflightName;
)

const (
    OfflinePreflight_NotRequired = C.kOfflinePreflight_NotRequired;
    OfflinePreflight_Optional = C.kOfflinePreflight_Optional;
    OfflinePreflight_Required = C.kOfflinePreflight_Required;
)

const (
    AudioUnitProperty_DistanceAttenuationData = C.kAudioUnitProperty_DistanceAttenuationData;
)

const (
    AudioUnitMigrateProperty_FromPlugin = C.kAudioUnitMigrateProperty_FromPlugin;
    AudioUnitMigrateProperty_OldAutomation = C.kAudioUnitMigrateProperty_OldAutomation;
)

const (
    OtherPluginFormat_Undefined = C.kOtherPluginFormat_Undefined;
    OtherPluginFormat_kMAS = C.kOtherPluginFormat_kMAS;
    OtherPluginFormat_kVST = C.kOtherPluginFormat_kVST;
    OtherPluginFormat_AU = C.kOtherPluginFormat_AU;
)

const (
    AudioUnitProperty_SampleRateConverterComplexity = C.kAudioUnitProperty_SampleRateConverterComplexity;
)

const (
    AudioUnitSampleRateConverterComplexity_Linear = C.kAudioUnitSampleRateConverterComplexity_Linear;
    AudioUnitSampleRateConverterComplexity_Normal = C.kAudioUnitSampleRateConverterComplexity_Normal;
    AudioUnitSampleRateConverterComplexity_Mastering = C.kAudioUnitSampleRateConverterComplexity_Mastering;
)

const (
    AudioOutputUnitProperty_CurrentDevice = C.kAudioOutputUnitProperty_CurrentDevice;
    AudioOutputUnitProperty_ChannelMap = C.kAudioOutputUnitProperty_ChannelMap;
    AudioOutputUnitProperty_EnableIO = C.kAudioOutputUnitProperty_EnableIO;
    AudioOutputUnitProperty_StartTime = C.kAudioOutputUnitProperty_StartTime;
    AudioOutputUnitProperty_SetInputCallback = C.kAudioOutputUnitProperty_SetInputCallback;
    AudioOutputUnitProperty_HasIO = C.kAudioOutputUnitProperty_HasIO;
    AudioOutputUnitProperty_StartTimestampsAtZero = C.kAudioOutputUnitProperty_StartTimestampsAtZero;
)

const (
    AUVoiceIOProperty_BypassVoiceProcessing = C.kAUVoiceIOProperty_BypassVoiceProcessing;
    AUVoiceIOProperty_VoiceProcessingEnableAGC = C.kAUVoiceIOProperty_VoiceProcessingEnableAGC;
    AUVoiceIOProperty_MuteOutput = C.kAUVoiceIOProperty_MuteOutput;
)

const (
    AUVoiceIOProperty_VoiceProcessingQuality = C.kAUVoiceIOProperty_VoiceProcessingQuality;
)

const (
    AUVoiceIOErr_UnexpectedNumberOfInputChannels = C.kAUVoiceIOErr_UnexpectedNumberOfInputChannels;
)

const (
    AudioUnitProperty_MeteringMode = C.kAudioUnitProperty_MeteringMode;
    AudioUnitProperty_MatrixLevels = C.kAudioUnitProperty_MatrixLevels;
    AudioUnitProperty_MatrixDimensions = C.kAudioUnitProperty_MatrixDimensions;
    AudioUnitProperty_MeterClipping = C.kAudioUnitProperty_MeterClipping;
)

const (
    AudioUnitProperty_3DMixerDistanceParams = C.kAudioUnitProperty_3DMixerDistanceParams;
    AudioUnitProperty_3DMixerAttenuationCurve = C.kAudioUnitProperty_3DMixerAttenuationCurve;
    AudioUnitProperty_SpatializationAlgorithm = C.kAudioUnitProperty_SpatializationAlgorithm;
    AudioUnitProperty_DopplerShift = C.kAudioUnitProperty_DopplerShift;
    AudioUnitProperty_3DMixerRenderingFlags = C.kAudioUnitProperty_3DMixerRenderingFlags;
    AudioUnitProperty_3DMixerDistanceAtten = C.kAudioUnitProperty_3DMixerDistanceAtten;
    AudioUnitProperty_ReverbPreset = C.kAudioUnitProperty_ReverbPreset;
)

const (
    K3DMixerAttenuationCurve_Power = C.k3DMixerAttenuationCurve_Power;
    K3DMixerAttenuationCurve_Exponential = C.k3DMixerAttenuationCurve_Exponential;
    K3DMixerAttenuationCurve_Inverse = C.k3DMixerAttenuationCurve_Inverse;
    K3DMixerAttenuationCurve_Linear = C.k3DMixerAttenuationCurve_Linear;
)

const (
    SpatializationAlgorithm_EqualPowerPanning = C.kSpatializationAlgorithm_EqualPowerPanning;
    SpatializationAlgorithm_SphericalHead = C.kSpatializationAlgorithm_SphericalHead;
    SpatializationAlgorithm_HRTF = C.kSpatializationAlgorithm_HRTF;
    SpatializationAlgorithm_SoundField = C.kSpatializationAlgorithm_SoundField;
    SpatializationAlgorithm_VectorBasedPanning = C.kSpatializationAlgorithm_VectorBasedPanning;
    SpatializationAlgorithm_StereoPassThrough = C.kSpatializationAlgorithm_StereoPassThrough;
)

const (
    K3DMixerRenderingFlags_InterAuralDelay = C.k3DMixerRenderingFlags_InterAuralDelay;
    K3DMixerRenderingFlags_DopplerShift = C.k3DMixerRenderingFlags_DopplerShift;
    K3DMixerRenderingFlags_DistanceAttenuation = C.k3DMixerRenderingFlags_DistanceAttenuation;
    K3DMixerRenderingFlags_DistanceFilter = C.k3DMixerRenderingFlags_DistanceFilter;
    K3DMixerRenderingFlags_DistanceDiffusion = C.k3DMixerRenderingFlags_DistanceDiffusion;
    K3DMixerRenderingFlags_LinearDistanceAttenuation = C.k3DMixerRenderingFlags_LinearDistanceAttenuation;
    K3DMixerRenderingFlags_ConstantReverbBlend = C.k3DMixerRenderingFlags_ConstantReverbBlend;
)

const (
    AudioUnitProperty_ScheduleAudioSlice = C.kAudioUnitProperty_ScheduleAudioSlice;
    AudioUnitProperty_ScheduleStartTimeStamp = C.kAudioUnitProperty_ScheduleStartTimeStamp;
    AudioUnitProperty_CurrentPlayTime = C.kAudioUnitProperty_CurrentPlayTime;
)

const (
    ScheduledAudioSliceFlag_Complete = C.kScheduledAudioSliceFlag_Complete;
    ScheduledAudioSliceFlag_BeganToRender = C.kScheduledAudioSliceFlag_BeganToRender;
    ScheduledAudioSliceFlag_BeganToRenderLate = C.kScheduledAudioSliceFlag_BeganToRenderLate;
)

const (
    AudioUnitProperty_ScheduledFileIDs = C.kAudioUnitProperty_ScheduledFileIDs;
    AudioUnitProperty_ScheduledFileRegion = C.kAudioUnitProperty_ScheduledFileRegion;
    AudioUnitProperty_ScheduledFilePrime = C.kAudioUnitProperty_ScheduledFilePrime;
    AudioUnitProperty_ScheduledFileBufferSizeFrames = C.kAudioUnitProperty_ScheduledFileBufferSizeFrames;
    AudioUnitProperty_ScheduledFileNumberBuffers = C.kAudioUnitProperty_ScheduledFileNumberBuffers;
)

const (
    AudioUnitProperty_ReverbRoomType = C.kAudioUnitProperty_ReverbRoomType;
    AudioUnitProperty_UsesInternalReverb = C.kAudioUnitProperty_UsesInternalReverb;
    MusicDeviceProperty_InstrumentName = C.kMusicDeviceProperty_InstrumentName;
    MusicDeviceProperty_InstrumentNumber = C.kMusicDeviceProperty_InstrumentNumber;
    MusicDeviceProperty_UsesInternalReverb = C.kMusicDeviceProperty_UsesInternalReverb;
    MusicDeviceProperty_BankName = C.kMusicDeviceProperty_BankName;
    MusicDeviceProperty_SoundBankData = C.kMusicDeviceProperty_SoundBankData;
    MusicDeviceProperty_StreamFromDisk = C.kMusicDeviceProperty_StreamFromDisk;
    MusicDeviceProperty_SoundBankFSRef = C.kMusicDeviceProperty_SoundBankFSRef;
    MusicDeviceProperty_SoundBankURL = C.kMusicDeviceProperty_SoundBankURL;
)

const (
    ReverbRoomType_SmallRoom = C.kReverbRoomType_SmallRoom;
    ReverbRoomType_MediumRoom = C.kReverbRoomType_MediumRoom;
    ReverbRoomType_LargeRoom = C.kReverbRoomType_LargeRoom;
    ReverbRoomType_MediumHall = C.kReverbRoomType_MediumHall;
    ReverbRoomType_LargeHall = C.kReverbRoomType_LargeHall;
    ReverbRoomType_Plate = C.kReverbRoomType_Plate;
    ReverbRoomType_MediumChamber = C.kReverbRoomType_MediumChamber;
    ReverbRoomType_LargeChamber = C.kReverbRoomType_LargeChamber;
    ReverbRoomType_Cathedral = C.kReverbRoomType_Cathedral;
    ReverbRoomType_LargeRoom2 = C.kReverbRoomType_LargeRoom2;
    ReverbRoomType_MediumHall2 = C.kReverbRoomType_MediumHall2;
    ReverbRoomType_MediumHall3 = C.kReverbRoomType_MediumHall3;
    ReverbRoomType_LargeHall2 = C.kReverbRoomType_LargeHall2;
)

const (
    AUSamplerProperty_LoadInstrument = C.kAUSamplerProperty_LoadInstrument;
    AUSamplerProperty_LoadAudioFiles = C.kAUSamplerProperty_LoadAudioFiles;
)

const (
    InstrumentType_DLSPreset = C.kInstrumentType_DLSPreset;
    InstrumentType_SF2Preset = C.kInstrumentType_SF2Preset;
    InstrumentType_AUPreset = C.kInstrumentType_AUPreset;
    InstrumentType_Audiofile = C.kInstrumentType_Audiofile;
    InstrumentType_EXS24 = C.kInstrumentType_EXS24;
)

const (
    AUSampler_DefaultPercussionBankMSB = C.kAUSampler_DefaultPercussionBankMSB;
    AUSampler_DefaultMelodicBankMSB = C.kAUSampler_DefaultMelodicBankMSB;
    AUSampler_DefaultBankLSB = C.kAUSampler_DefaultBankLSB;
)

const (
    AudioUnitProperty_DeferredRendererPullSize = C.kAudioUnitProperty_DeferredRendererPullSize;
    AudioUnitProperty_DeferredRendererExtraLatency = C.kAudioUnitProperty_DeferredRendererExtraLatency;
    AudioUnitProperty_DeferredRendererWaitFrames = C.kAudioUnitProperty_DeferredRendererWaitFrames;
)

const (
    AUNetReceiveProperty_Hostname = C.kAUNetReceiveProperty_Hostname;
    AUNetReceiveProperty_Password = C.kAUNetReceiveProperty_Password;
)

const (
    AUNetSendProperty_PortNum = C.kAUNetSendProperty_PortNum;
    AUNetSendProperty_TransmissionFormat = C.kAUNetSendProperty_TransmissionFormat;
    AUNetSendProperty_TransmissionFormatIndex = C.kAUNetSendProperty_TransmissionFormatIndex;
    AUNetSendProperty_ServiceName = C.kAUNetSendProperty_ServiceName;
    AUNetSendProperty_Disconnect = C.kAUNetSendProperty_Disconnect;
    AUNetSendProperty_Password = C.kAUNetSendProperty_Password;
)

const (
    AUNetSendPresetFormat_PCMFloat32 = C.kAUNetSendPresetFormat_PCMFloat32;
    AUNetSendPresetFormat_PCMInt24 = C.kAUNetSendPresetFormat_PCMInt24;
    AUNetSendPresetFormat_PCMInt16 = C.kAUNetSendPresetFormat_PCMInt16;
    AUNetSendPresetFormat_Lossless24 = C.kAUNetSendPresetFormat_Lossless24;
    AUNetSendPresetFormat_Lossless16 = C.kAUNetSendPresetFormat_Lossless16;
    AUNetSendPresetFormat_ULaw = C.kAUNetSendPresetFormat_ULaw;
    AUNetSendPresetFormat_IMA4 = C.kAUNetSendPresetFormat_IMA4;
    AUNetSendPresetFormat_AAC_128kbpspc = C.kAUNetSendPresetFormat_AAC_128kbpspc;
    AUNetSendPresetFormat_AAC_96kbpspc = C.kAUNetSendPresetFormat_AAC_96kbpspc;
    AUNetSendPresetFormat_AAC_80kbpspc = C.kAUNetSendPresetFormat_AAC_80kbpspc;
    AUNetSendPresetFormat_AAC_64kbpspc = C.kAUNetSendPresetFormat_AAC_64kbpspc;
    AUNetSendPresetFormat_AAC_48kbpspc = C.kAUNetSendPresetFormat_AAC_48kbpspc;
    AUNetSendPresetFormat_AAC_40kbpspc = C.kAUNetSendPresetFormat_AAC_40kbpspc;
    AUNetSendPresetFormat_AAC_32kbpspc = C.kAUNetSendPresetFormat_AAC_32kbpspc;
    AUNetSendPresetFormat_AAC_LD_64kbpspc = C.kAUNetSendPresetFormat_AAC_LD_64kbpspc;
    AUNetSendPresetFormat_AAC_LD_48kbpspc = C.kAUNetSendPresetFormat_AAC_LD_48kbpspc;
    AUNetSendPresetFormat_AAC_LD_40kbpspc = C.kAUNetSendPresetFormat_AAC_LD_40kbpspc;
    AUNetSendPresetFormat_AAC_LD_32kbpspc = C.kAUNetSendPresetFormat_AAC_LD_32kbpspc;
    AUNetSendNumPresetFormats = C.kAUNetSendNumPresetFormats;
)

const (
    NilAudioComponent = AudioComponent(0);
)

type AudioComponent uintptr

type AudioComponentInstance uintptr

type AudioUnit AudioComponentInstance

type OSStatus C.OSStatus

func build_AudioComponentDescription(x *AudioComponentDescription) *C.AudioComponentDescription {
    return &C.AudioComponentDescription {
        C.OSType(x.ComponentType),
        C.OSType(x.ComponentSubType),
        C.OSType(x.ComponentManufacturer),
        C.UInt32(x.ComponentFlags),
        C.UInt32(x.ComponentFlagsMask),
    }
}

func AudioComponentCount(desc AudioComponentDescription) uint32 {
    return uint32(C.AudioComponentCount(build_AudioComponentDescription(&desc)))
}

func AudioComponentFindNext(component AudioComponent, desc AudioComponentDescription) AudioComponent {
    return AudioComponent(unsafe.Pointer(C.AudioComponentFindNext(C.AudioComponent(unsafe.Pointer(component)), build_AudioComponentDescription(&desc))))
}

func AudioComponentInstanceNew(component AudioComponent) (AudioComponentInstance, OSStatus) {
    var x C.AudioComponentInstance
    status := C.AudioComponentInstanceNew(C.AudioComponent(unsafe.Pointer(component)), &x)
    return AudioComponentInstance(unsafe.Pointer(x)), OSStatus(status)
}

func AudioUnitInitialize(unit AudioUnit) OSStatus {
    return OSStatus(C.AudioUnitInitialize(C.AudioUnit(unsafe.Pointer(unit))))
}

func AudioUnitUninitialize(unit AudioUnit) OSStatus {
    return OSStatus(C.AudioUnitUninitialize(C.AudioUnit(unsafe.Pointer(unit))))
}

func AudioOutputUnitStart(unit AudioUnit) OSStatus {
    return OSStatus(C.AudioOutputUnitStart(C.AudioUnit(unsafe.Pointer(unit))))
}

func AudioOutputUnitStop(unit AudioUnit) OSStatus {
    return OSStatus(C.AudioOutputUnitStop(C.AudioUnit(unsafe.Pointer(unit))))
}

func AudioUnitGetPropertyInfo(unit AudioUnit, id uint32, scope uint32, element uint32) (uintptr, bool, OSStatus) {
    var _dataSize C.UInt32
    var _writable C.Boolean
    status := C.AudioUnitGetPropertyInfo(
        C.AudioUnit(unsafe.Pointer(unit)),
        C.AudioUnitPropertyID(id),
        C.AudioUnitScope(scope),
        C.AudioUnitElement(element),
        &_dataSize,
        &_writable,
    )
    return uintptr(_dataSize), _writable != 0, OSStatus(status)
}

func AudioUnitGetProperty(unit AudioUnit, id uint32, scope uint32, element uint32) (interface {}, OSStatus) {
    _dataSize, _, _status := AudioUnitGetPropertyInfo(unit, id, scope, element)
    if _status != 0 {
        return nil, _status
    }
    dataSize := C.UInt32(_dataSize)
    retval := make([]byte, uint32(dataSize))
    status := OSStatus(C.AudioUnitGetProperty(
        C.AudioUnit(unsafe.Pointer(unit)),
        C.AudioUnitPropertyID(id),
        C.AudioUnitScope(scope),
        C.AudioUnitElement(element),
        unsafe.Pointer(&retval[0]),
        &dataSize,
    ))
    return retval, status
}

func AudioUnitGetProperty_StreamFormat(unit AudioUnit, scope uint32, element uint32) (AudioStreamBasicDescription, OSStatus) {
    _data := C.AudioStreamBasicDescription {}
    _dataSize := reflect.ValueOf(_data).Type().Size()
    data := unsafe.Pointer(&_data)
    dataSize := C.UInt32(_dataSize)
    status := OSStatus(C.AudioUnitGetProperty(
        C.AudioUnit(unsafe.Pointer(unit)),
        AudioUnitProperty_StreamFormat,
        C.AudioUnitScope(scope),
        C.AudioUnitElement(element),
        unsafe.Pointer(data),
        &dataSize,
    ))
    retval := AudioStreamBasicDescription {
        float64(_data.mSampleRate),
        uint32(_data.mFormatID),
        uint32(_data.mFormatFlags),
        uint32(_data.mBytesPerPacket),
        uint32(_data.mFramesPerPacket),
        uint32(_data.mBytesPerFrame),
        uint32(_data.mChannelsPerFrame),
        uint32(_data.mBitsPerChannel),
    }
    return retval, status
}

func AudioUnitSetProperty(unit AudioUnit, id int, scope int, element int, data interface {}) OSStatus {
    rv := reflect.ValueOf(data)
    var rawDataSize uintptr
    var rawData unsafe.Pointer
    switch t := data.(type) {
    case AudioStreamBasicDescription:
        _data := build_AudioStreamBasicDescription(t)
        rv = reflect.ValueOf(_data)
        rawDataSize = rv.Type().Size()
        rawData = unsafe.Pointer(rv.UnsafeAddr())
    case AudioObjectID:
        _data := C.AudioObjectID(t)
        rawDataSize = 4 // XXX
        rawData = unsafe.Pointer(&_data)
    case bool:
        var _data C.UInt32
        if t {
            _data = C.UInt32(1)
        } else {
            _data = C.UInt32(0)
        }
        rawDataSize = 4 // XXX
        rawData = unsafe.Pointer(&_data)
    default:
        rawDataSize = rv.Type().Elem().Size() * uintptr(rv.Len())
        // assert rv.Type().Elem().Size() == rv.Type().Elem().Align()
        rawData = unsafe.Pointer(rv.Index(0).UnsafeAddr())
    }
    return OSStatus(C.AudioUnitSetProperty(
        C.AudioUnit(unsafe.Pointer(unit)),
        C.AudioUnitPropertyID(id),
        C.AudioUnitScope(scope),
        C.AudioUnitElement(element),
        rawData,
        C.UInt32(rawDataSize),
    ))
}

