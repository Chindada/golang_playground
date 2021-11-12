```go
package network

import (
 "fmt"
 "moxa/mxview/pkg/core"
 "moxa/mxview/pkg/core/context"
 "moxa/mxview/pkg/core/eventbus/topic"
)


func setTurboRingV2(deviceID int) {
 ctx := context.NewContext()
 ctx.Set("id", deviceID)
 e := context.Engine{}
 e.Use(updateTRv2BlockingPort)
 e.Use(updateTRv2RedundancyRole)
 e.Run(ctx)
}

func updateTRv2BlockingPort(ctx context.Context) {
 var linksToPublish []*core.LinkData
 id := ctx.GetInt("id")
 links, _ := agent.links.GetConnectedLinksByDevice(id)

 // turboRingV2CouplingEnableStatus, _ := cache.Get("device:%d:turboRingV2CouplingEnable:0", id)
 for _, link := range links {
  if id == link.FromDeviceID {
   ring1StatusMapArr := getTurboRingV2PortStatus(id)
   if len(ring1StatusMapArr) == 0 {
    continue
   }
   ring2StatusMapArr := getTurboRingV2PortStatus(link.ToDeviceID)
   if len(ring2StatusMapArr) == 0 {
    continue
   }
   if ring1StartStatus, ok := ring1StatusMapArr[0][link.FromPort]; ok {
    if ring1EndStatus, ok := ring2StatusMapArr[0][link.ToPort]; ok {
     startTurboRingV2Ring1Enable, _ := cache.Get("device:%d:turboRingV2Ring1Enable:0", id)
     endTurboRingV2Ring1Enable, _ := cache.Get("device:%d:turboRingV2Ring1Enable:0", link.ToDeviceID)
     if startTurboRingV2Ring1Enable == endTurboRingV2Ring1Enable && endTurboRingV2Ring1Enable == turboRingV2Enable {
      link.RedundancyType = redundancyProtocolTubroRing2
     }
     if ring1StartStatus == ring1EndStatus && ring1EndStatus == turboRingV2Blocking {
      fmt.Println(id)
      fmt.Println(ring1EndStatus)
      fmt.Println(ring1StartStatus)
      fmt.Println(id)
      link.IsRedundancyBlocking = true
     }
    }
    if ring2EndStatus, ok := ring2StatusMapArr[1][link.ToPort]; ok {
     startTurboRingV2Ring1Enable, _ := cache.Get("device:%d:turboRingV2Ring1Enable:0", id)
     endTurboRingV2Ring2Enable, _ := cache.Get("device:%d:turboRingV2Ring2Enable:0", link.ToDeviceID)
     if startTurboRingV2Ring1Enable == endTurboRingV2Ring2Enable && endTurboRingV2Ring2Enable == turboRingV2Enable {
      link.RedundancyType = redundancyProtocolTubroRing2
     }
     if ring1StartStatus == ring2EndStatus && ring2EndStatus == turboRingV2Blocking {
      fmt.Println(id)
      fmt.Println(ring1StartStatus)
      fmt.Println(ring2EndStatus)
      fmt.Println(id)
      link.IsRedundancyBlocking = true
     }
    }
   }

   if ring2StartStatus, ok := ring1StatusMapArr[1][link.FromPort]; ok {
    if ring1EndStatus, ok := ring2StatusMapArr[0][link.ToPort]; ok {
     startTurboRingV2Ring1Enable, _ := cache.Get("device:%d:turboRingV2Ring1Enable:0", id)
     endTurboRingV2Ring1Enable, _ := cache.Get("device:%d:turboRingV2Ring1Enable:0", link.ToDeviceID)
     if startTurboRingV2Ring1Enable == endTurboRingV2Ring1Enable && endTurboRingV2Ring1Enable == turboRingV2Enable {
      link.RedundancyType = redundancyProtocolTubroRing2
     }
     if ring2StartStatus == ring1EndStatus && ring1EndStatus == turboRingV2Blocking {
      link.IsRedundancyBlocking = true
     }
    }
    if ring2EndStatus, ok := ring2StatusMapArr[1][link.ToPort]; ok {
     startTurboRingV2Ring1Enable, _ := cache.Get("device:%d:turboRingV2Ring1Enable:0", id)
     endTurboRingV2Ring2Enable, _ := cache.Get("device:%d:turboRingV2Ring2Enable:0", link.ToDeviceID)
     if startTurboRingV2Ring1Enable == endTurboRingV2Ring2Enable && endTurboRingV2Ring2Enable == turboRingV2Enable {
      link.RedundancyType = redundancyProtocolTubroRing2
     }
     if ring2StartStatus == ring2EndStatus && ring2EndStatus == turboRingV2Blocking {
      link.IsRedundancyBlocking = true
     }
    }
   }

   if couplingStartStatus, ok := getTurboRingV2BlockingStatus(id)[link.FromPort]; ok {
    if couplingEndStatus, ok := getTurboRingV2BlockingStatus(link.ToDeviceID)[link.ToPort]; ok {
     startTurboRingV2CouplingEnableStatus, _ := cache.Get("device:%d:turboRingV2CouplingEnable:0", id)
     endTurboRingV2CouplingEnableStatus, _ := cache.Get("device:%d:turboRingV2CouplingEnable:0", link.ToDeviceID)
     if startTurboRingV2CouplingEnableStatus == endTurboRingV2CouplingEnableStatus && endTurboRingV2CouplingEnableStatus == turboRingV2Enable {
      link.RedundancyType = redundancyProtocolTubroRing2
     }
     if couplingStartStatus == couplingEndStatus && couplingEndStatus == turboRingV2Blocking {
      link.IsRedundancyBlocking = true
     }
    }
   }
   linksToPublish = append(linksToPublish, link)
  }
 }
 for _, link := range linksToPublish {
  bus.Publish(topic.LinkUpdate, *link)
 }
 ctx.Next()
}

func updateTRv2RedundancyRole(ctx context.Context) {
 id := ctx.GetInt("id")

 turboRingV2Ring1Master, _ := cache.Get("device:%d:turboRingV2Ring1Master:0", id)
 turboRingV2Ring2Master, _ := cache.Get("device:%d:turboRingV2Ring2Master:0", id)

 if turboRingV2Ring1Master == turboRingV2MasterYes {
  bus.Publish(topic.DeviceUpdateRedundancy, id, deviceTrv2Ring1Master)
 } else if turboRingV2Ring2Master == turboRingV2MasterYes {
  bus.Publish(topic.DeviceUpdateRedundancy, id, deviceTrv2Ring2Master)
 }
}

func getTurboRingV2PortStatus(id int) (portStatusMapArr []map[int]string) {
 turboRingV2Ring1Port1, _ := cache.Get("device:%d:turboRingV2Ring1Port1:0", id)
 turboRingV2Ring1Port2, _ := cache.Get("device:%d:turboRingV2Ring1Port2:0", id)
 turboRingV2Ring2Port1, _ := cache.Get("device:%d:turboRingV2Ring2Port1:0", id)
 turboRingV2Ring2Port2, _ := cache.Get("device:%d:turboRingV2Ring2Port2:0", id)
 turboRingV2Ring1Port1Status, _ := cache.Get("device:%d:turboRingV2Ring1Port1Status:0", id)
 turboRingV2Ring1Port2Status, _ := cache.Get("device:%d:turboRingV2Ring1Port2Status:0", id)
 turboRingV2Ring2Port1Status, _ := cache.Get("device:%d:turboRingV2Ring2Port1Status:0", id)
 turboRingV2Ring2Port2Status, _ := cache.Get("device:%d:turboRingV2Ring2Port2Status:0", id)

 ring1PortRDTMap := make(map[int]string)
 var ring1Port1, ring1Port2 int
 var ok bool
 if ring1Port1, ok = turboRingV2Ring1Port1.(int); !ok {
  return portStatusMapArr
 }
 ring1PortRDTMap[ring1Port1] = turboRingV2Ring1Port1Status.(string)
 if ring1Port2, ok = turboRingV2Ring1Port2.(int); !ok {
  return portStatusMapArr
 }
 ring1PortRDTMap[ring1Port2] = turboRingV2Ring1Port2Status.(string)

 ring2PortRDTMap := make(map[int]string)
 var ring2Port1, ring2Port2 int
 if ring2Port1, ok = turboRingV2Ring2Port1.(int); !ok {
  return portStatusMapArr
 }
 ring2PortRDTMap[ring2Port1] = turboRingV2Ring2Port1Status.(string)
 if ring2Port2, ok = turboRingV2Ring2Port2.(int); !ok {
  return portStatusMapArr
 }
 ring2PortRDTMap[ring2Port2] = turboRingV2Ring2Port2Status.(string)
 return []map[int]string{ring1PortRDTMap, ring2PortRDTMap}
}

func getTurboRingV2BlockingStatus(id int) map[int]string {
 turboRingV2CouplingPort1, _ := cache.Get("device:%d:turboRingV2CouplingPort1:0", id)
 turboRingV2CouplingPort1Status, _ := cache.Get("device:%d:turboRingV2CouplingPort1Status:0", id)
 turboRingV2CouplingPort2, _ := cache.Get("device:%d:turboRingV2CouplingPort2:0", id)
 turboRingV2CouplingPort2Status, _ := cache.Get("device:%d:turboRingV2CouplingPort2Status:0", id)

 couplingPortMap := make(map[int]string)
 couplingPortMap[turboRingV2CouplingPort1.(int)] = turboRingV2CouplingPort1Status.(string)
 couplingPortMap[turboRingV2CouplingPort2.(int)] = turboRingV2CouplingPort2Status.(string)

 return couplingPortMap
}
```
