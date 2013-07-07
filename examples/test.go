package main

import "github.com/moriyoshi/go-coreaudio/coreaudio"

func main() {
    desc := coreaudio.AudioComponentDescription {
        ComponentType: coreaudio.AudioUnitType_Output,
        ComponentSubType: coreaudio.AudioUnitSubType_HALOutput,
        ComponentManufacturer: coreaudio.AudioUnitManufacturer_Apple,
        ComponentFlags: 0,
        ComponentFlagsMask: 0,
    }
    print(coreaudio.AudioComponentCount(desc), "\n")

    comp := coreaudio.AudioComponentFindNext(coreaudio.AudioComponent(coreaudio.NilAudioComponent), desc)
    if comp == coreaudio.NilAudioComponent {
        return
    }

    print(comp, "\n")
    au, err := coreaudio.AudioComponentInstanceNew(comp)
    print(au, ", err=", err, "\n")

    format, status := coreaudio.AudioUnitGetProperty_StreamFormat(
        coreaudio.AudioUnit(au),
        coreaudio.AudioUnitScope_Input,
        1,
    )

    if status != 0 {
        print("err=", status, "\n")
        return
    }

    print("sampleRate=", format.SampleRate, "\n")
    print("formatId=", format.FormatID, "\n")
    print("formatFlags=", format.FormatFlags, "\n")
    print("bytesPerPacket=", format.BytesPerPacket, "\n")
    print("framesPerPacket=", format.FramesPerPacket, "\n")
    print("bytesPerFrame=", format.BytesPerFrame, "\n")
    print("channelsPerFrame=", format.ChannelsPerFrame, "\n")
    print("bitsPerChannel=", format.BitsPerChannel, "\n")

    coreaudio.AudioUnitSetProperty(
        coreaudio.AudioUnit(au),
        coreaudio.AudioOutputUnitProperty_EnableIO,
        coreaudio.AudioUnitScope_Output,
        0,
        true,
    )

    coreaudio.AudioUnitSetProperty(
        coreaudio.AudioUnit(au),
        coreaudio.AudioOutputUnitProperty_EnableIO,
        coreaudio.AudioUnitScope_Input,
        1,
        true,
    )

    inputDevice, err := coreaudio.AudioObjectGetPropertyData_SystemObject(
        coreaudio.AudioObjectPropertyAddress {
            coreaudio.AudioHardwarePropertyDefaultOutputDevice,
            coreaudio.AudioObjectPropertyScopeGlobal,
            0,
        },
    )

    print("err=", err, "\n")

    if err != 0 {
        return
    }

    err = coreaudio.AudioUnitSetProperty(
        coreaudio.AudioUnit(au),
        coreaudio.AudioOutputUnitProperty_CurrentDevice,
        coreaudio.AudioUnitScope_Global,
        0,
        inputDevice,
    )

    print("err=", err, "\n")
}
