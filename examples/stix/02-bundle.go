// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"encoding/json"
	"fmt"

	"github.com/freetaxii/libstix2/objects"
)

func main() {
	sm := objects.NewBundle()

	// Create a campaign
	c := objects.NewCampaign("2.0")
	c.SetName("Bank Attack 2016")
	c.SetObjective("Compromise SWIFT system and steal money")
	sm.AddObject(c)

	// Create an indicator
	i := objects.NewIndicator("2.0")
	i.SetName("Malware C2 Indicator 2016")
	i.SetDescription("This indicator should detect the SpyEye malware by looking for this MD5 hash")
	i.SetPattern("file-object:hashes.md5 = 84714c100d2dfc88629531f6456b8276")
	sm.AddObject(i)

	// Define some infrastructure as used by this campaign and malware.
	infra := objects.NewInfrastructure("2.0")
	infra.SetName("SpyEye Command and Control Servers")
	infra.SetDescription("These servers are located in a datacenter in the Netherlands and the IPs change on a weekly basis")
	// infra.AddKillChainPhase("lockheed-martin-cyber-kill-chain", "command-and-control")
	// infra.SetFirstSeenText("2016-09-01T00:00:01Z")
	// infra.SetRegion("Europe")
	// infra.SetCountry("NL")
	sm.AddObject(infra)

	// Define some Observed Data for the Infrastructure
	od1 := objects.NewObservedData("2.0")
	od1.SetFirstObserved("2016-09-01T00:00:01Z")
	od1.SetLastObserved("2016-09-07T00:00:01Z")
	od1.SetNumberObserved(3)
	//od1.SetCybox("This will be a CybOX container object using the ipv4-addr object pointing to 5.79.68.0/24")
	sm.AddObject(od1)

	od2 := objects.NewObservedData("2.0")
	od2.SetFirstObserved("2016-09-07T00:00:01Z")
	od2.SetLastObserved("2016-09-14T00:00:01Z")
	od2.SetNumberObserved(3)
	//od2.SetCybox("This will be a CybOX container object using the ipv4-addr object pointing to 5.79.52.0/24")
	sm.AddObject(od2)

	// Define some Observed Data for the sighting of the Infrastructure
	od3 := objects.NewObservedData("2.0")
	od3.SetFirstObserved("2016-09-07T00:00:01Z")
	od3.SetLastObserved("2016-09-14T00:00:01Z")
	od3.SetNumberObserved(1)
	//od3.SetCybox("This will be a CybOX container object using the ipv4-addr object pointing to 5.79.52.100")
	sm.AddObject(od3)

	// Define a family of malware
	m1 := objects.NewMalware("2.0")
	m1.SetName("Zeus")
	m1.AddLabel("trojan")
	m1.AddLabel("malware-family")
	sm.AddObject(m1)

	// Define a piece of malware
	m2 := objects.NewMalware("2.0")
	m2.SetName("SpyEye")
	m2.AddLabel("trojan")
	// m2.AddFilename("cleansweep.exe")
	// m2.AddFilename("spyeye2_exe")
	// m2.AddFilename("build_1_.exe")
	// m2.AddHash("md5", "84714c100d2dfc88629531f6456b8276")
	// m2.AddHash("sha256", "861aa9c5ddcb5284e1ba4e5d7ebacfa297567c353446506ee4b4e39c84454b09")
	// m2.AddKillChainPhase("lockheed-martin-cyber-kill-chain", "command-and-control")
	sm.AddObject(m2)

	// Define some scan data for the malware sample
	// m2s1 := m2.NewScanData()
	// m2s1.SetScannedText("2016-08-30T06:31:48Z")
	// m2s1.SetProduct("avg")
	// m2s1.SetClassification("Generic16.BFGI")

	// m2s2 := m2.NewScanData()
	// m2s2.SetScannedText("2016-08-30T06:31:48Z")
	// m2s2.SetProduct("avast")
	// m2s2.SetClassification("Win32:Downloader-NTU [PUP]")

	// Connect the malware sample to a malware family
	r1 := objects.NewRelationship("2.0")
	r1.SetRelationshipType("member-of")
	r1.SetSourceTarget(m1.GetID(), m2.GetID())
	sm.AddObject(r1)

	// Identify that this campaign uses this piece of malware
	r2 := objects.NewRelationship("2.0")
	r2.SetRelationshipType("uses")
	r2.SetSourceTarget(c.GetID(), m2.GetID())
	sm.AddObject(r2)

	// Identify that this campaign uses this infrastructure
	r3 := objects.NewRelationship("2.0")
	r3.SetRelationshipType("uses")
	r3.SetSourceTarget(c.GetID(), infra.GetID())
	sm.AddObject(r3)

	// Identify that this malware uses this infrastructure
	r4 := objects.NewRelationship("2.0")
	r4.SetRelationshipType("uses")
	r4.SetSourceTarget(m2.GetID(), infra.GetID())
	sm.AddObject(r4)

	// Identify that this indicator can indicate the presence of this malware
	r5 := objects.NewRelationship("2.0")
	r5.SetRelationshipType("indicates")
	r5.SetSourceTarget(i.GetID(), m2.GetID())
	sm.AddObject(r5)

	// Attach some Observed Data to an Infrastructure Object
	r6 := objects.NewRelationship("2.0")
	r6.SetRelationshipType("part-of")
	r6.SetSourceTarget(od1.GetID(), infra.GetID())
	sm.AddObject(r6)

	// Attach some Observed Data to an Infrastructure Object
	r7 := objects.NewRelationship("2.0")
	r7.SetRelationshipType("part-of")
	r7.SetSourceTarget(od2.GetID(), infra.GetID())
	sm.AddObject(r7)

	// Add a sighting for the malware
	s1 := objects.NewSighting("2.0")
	s1.SetFirstSeen("2016-09-01T00:00:01Z")
	s1.SetLastSeen("2016-09-01T10:30:00Z")
	s1.SetCount(3)
	s1.SetSightingOfRef(m2.GetID())
	sm.AddObject(s1)

	// Add a sighting for the infrastructure
	s2 := objects.NewSighting("2.0")
	s2.SetFirstSeen("2016-09-01T00:00:01Z")
	s2.SetLastSeen("2016-09-01T10:30:00Z")
	s2.SetCount(10)
	s2.SetSightingOfRef(infra.GetID())
	s2.AddObservedDataRef(od3.GetID())
	sm.AddObject(s2)

	var data []byte
	data, _ = json.MarshalIndent(sm, "", "    ")

	fmt.Println(string(data))
}
