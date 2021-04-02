/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify"
)

// rollCmd represents the roll command
var rollCmd = &cobra.Command{
	Use:   "roll",
	Short: "say hi to Rick",
	Long: "serenade yourself with the golden voice of Rick Astley",
	Run: func(cmd *cobra.Command, args []string) {
		roll()
	},
}

func init() {
	rootCmd.AddCommand(rollCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rollCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rollCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func roll() {
	client := getAuthenticatedClientWithRetry()
	searchResults, err := client.Search("never gonna give you up", spotify.SearchTypeTrack)
	check(err)

	opts := spotify.PlayOptions{
		DeviceID:        nil,
		PlaybackContext: nil,
		URIs:            nil,
		PlaybackOffset:  nil,
		PositionMs:      0,
	}

	uri := searchResults.Tracks.Tracks[0].URI
	opts.URIs = append(opts.URIs, uri)

	err = client.PlayOpt(&opts)
	check(err)

	fmt.Print(rick)
}

const rick = `
                                         ..,,,,,'.......
                                      .,:llllddolodxxxxoc;,.
                                     .cdkkkOOOOkkOOOkkkkkkxdc.
                                   .,codxkOO0OOOOOOOkkkOkkxxxc.
                                .';cloddkkOO00OOO0OOOkOOOOOkxxc.
                              .;cllloddxkkOOOOO0000OOOOOOOOkxxd;
                              '::cclodddxxxxxxkkkkkkkkkkkkkkkxxo.
                              ,:coooodooddooolloooollllooodxkkxx:
                             .,:oxkxdoodddoollllllccccccccloxkkxc.
                             .,lxkxddodddddooolllllccccccccldxkd,
                             .;oxkxddddddddooolllllllcccccclodxo.
                             .coxxdooddxxxxxxdoolllllcccccclodxo.
                             'coddoodddxxxdxxxxdoloddoollllloxxl.
                           ..:cllllloddddddddxxdolldxxxdoollddc.
                         .;clooc;;;coooooooodxddolcloddoolloo;
                         .:ooool::clodddoodddddollccclcccccoc.
                          .:lddolcloodddddddxxxdoollccccccclc.
                            'lodoolooodddddddddoolllllccllllc.
                             .:llooooooddddddoollcccllllllll;
                               ':loooodddxkkkxxxddollllllll;.
                               .;cooooodddxxxddxxxdlllll:,.
                               .;cloooooddddddoollclllc,.
                               .;llodooddoooooolllclll;.
                               .;oooddddddoolllcclllll;
                               .;loodddddddddoollllllll,
                               ':lodddddddddddooolllllcoo;..
                             .';:llllloooddoolllllllll:lxOkdl;'...
                        ..',,;;:clcccclllodoolllllllll:cxOO00Oxdlcc:,'...
                    ..';:ccc:;:lodoodddoodddooollllol:;cxO00000Oxdxxddollc:,'..
                ..';coooolcc:;cloddoooooodddddoooool::;lkO0000KK0Okkkkxxxxxxxxdoc;,...
            ..';codxkkkkkdl::::lodddoddooddddddddoc:;;;lk00000KKKK00OOOOOOOOOOOkkkxxdol;.
        ..';clodkOO00000Oxl::::codddooooodxxxxdoc:::;;cxO0000KKKKKKK0OOO000000000000000Ox;
     .';clodxkO0000000000kl:;:cldddddddoodxxdoc:::;:;cdO00000KKKKKKKKK0OO0000K000000000OOk:
   .;lodxxkO000000K000000kl;;:coxddddddoodolcc:::;::lxO000000KKKKKKKKK0000KK00000000000OOOx'
   .okkOO00000000KKK00000Ol::ccldooollllllcc::::::cok0000000KKKKKKKKKKKKKKKKK000000000000OOc
   'k00000000000KKKKKK000Odlooddddoolloooooolc:::cldk000000KKKKKKKKKKKKKKKKK00000000000000Oo.
   ,O00000000000KKKKKKK0000O00OkddoccodxxddxxdolodxkO000KKKKKKKKKKKKKK00KKKK000000000000000d.
   :O00000000000KKKKKKKK0000000Oxoc;;:odxxxxxdxxxxxkO00KKKKKKKKKKKKKKOxddk000000KKK00000000o.
   c00000000000KK00KKKKKK000K00kxoc:::lddddddxxxxxkO0000KKKKKKKKKKKKKOddooxO00KKK000000000x,
  .l000000000KKKKK000KKKK000000Oxoc:::odxxxxkkkkkkkO0000KKKKKKKKKKKKKOxddodOKKK00OOO0K000x'
  .l000000000KKKKK0000KKK000000OkoccccoddxxxxxxxxxxO0000KKKKKKKKKKKKK0kxxddOKK00kdloxO00O:
   :O000KKKKKKKKKKK000KK0000000Oxoc::codxxxxxxkkkkkO000KKKKKKKKKKKKKKKOxxdox0Okdolccldk0O:
   ,k000KKKKKKKKKKK00KKK0KKKKK00kdlccldxxxxxxxxxxxxk00KKKKKKKKKKKKKKKKOxddooxdolllodxkO0k;
   .d000KKKKKKKKKKKKKKKKKKKKKK00kdddoddddxxxxxxxxkkO00KKKKKKKKKKKKKKKK0xdxdddooloodkOOO0o.
   .o000000KKKKKKKKKKKKKKKKKKKK0OkxdxkkkkkkkkkkxkkkkO0KKKKKKKKKKKKKKKK0xdddddddllllxO0OOd'
   .o000000KKKKKKKKKKKKKKKKKKKK0OxdddxxxxxddddxxxxxkO0KKXKKKKKKKKKKKKK0xxxdooddoooookO000x,
   .o00000KKKKKKKKKKKKKKKKKKKKKK0kdddxkkkkkkkkkkkkkOO0KKKXXKXXKKXXKXKK0kxxdooolloooxO00O00Ol.
   .d00000KKKKKKKKKKKKKKKKKKKKKK0kddxxxxxxxxdddddxxxkO0KKXXXXXXXXKKXK00Oxxxdololllok0000000Ox,
   .d00000KKKKKKKKKKKKKKKKKKKKKKKkddxxxxxkkkkkkkkkkOO00KKKXXXXXXXXXX0kdxdddocclooodk00KKK0000x,
    c0000KKKKKKKKKKKKKKKKKKKKKKKKOxdxkkkkkkkxxxxxxxkkkO0KKXXXXXXXXXXKOxdolllcclllldkO000KK0000d.
    'k000KKKKKKKKKKKKKKKKKKKKKKKKOddxxxxxxxkkkkkkO0000000KKXXXXXXXXXXK0OxdoooooooooxO000000000Oo.
    .x000KKKKKKKKKKKKKKKKKKKKKKKKOxxkOkkkkkkkkkkxkkkkkkkO0KKXXXXXXXXXXXKK0kxxxddddxO00000000000Ol.
    .d00KKKKKKKKKKKKKKKKKKKKKKKKKOxxkkkxxxxxxxxxkkkkkkkkkO0KKKXXXXXXXXXXXXK000OOOO0KK00000000000O:
    .dK0KKKKKKKKKKKKKKKKKKKKKKKKKOxxkkkkkkkkkkkkkkkkOOOOO00KKKXXXXXXXXXXXXXKKKKK00KXKK000000000O0x'
    .xKKKKKKKKKKKKKKKKKKKKKKKKKKKOxxkkkkkkkkkkkkkkkkkkkOOOO0KKKXXXXXXXXXXXXXKKK000KXXKK0000000000O;
    'kKKKKKKKKKKKKKKKKKKKKKKKKKKKOxxkxxxkkkkkkkkkOOkOOOO00000KKKXXXXXXXXXXXXXXK000KXXXK0000000000O:
    'kKKKKKKXKKKKKKKKKKKKKKKKKKKKOxkOOkkOkkkkkkkkkkkkkkkOOOO00KKKKXXXXXXXXXXNXXK000XXXK0OO00000000o.
    .xKKKKKXXXKKKKKKKKKKKKKKKKKKKOxxkkxxxxkkkkkkkkkkkOOOOO0000KKKKKKXXXXXXXNNNXXXKKKXKK0OkO000000k;
    .dKKKKKKKXXKKKKKKKKKKKKKKKKKKOxkOkkkkkkkkkkkkkkkkkkkkkOO0000KKKKKKXXXXNNXXNXXXXKKKKK00000000k,
     c000KKKXXXKKKKKKKKKKKKKKKKKKOkkkkkxxxxxkkkkkkkkkkkkOOOOOOOOOKKKKKKXXXXXXNNNNXXK0KXXKKKKK0ko,
     ,xO00KXXXXXKKKKKKKKKKKKKKKKKOkkOOOOkkkkkkkkkkkkkkkkkOOO000OOOKKKKKXXXXXXXXXNXXx;,cok00ko:.
     ;xk0KKXXXXXXKKKKKKKKKKKKKKKKOkkkkkkxxxkkkkkkxxxxkkkOOOO00000OOKKKKXXXXXXXXXXXXKk;   ...
     :O0KKKKXXXXXKKKKKKKKKKKKKKKKOkkOOOOOOOOO000OOkkkkkOkkkOO0000kkOKKKKXXXKXXXKXXKKK0o.
     ,kKKKKKKXXXXXKKKKKKKKKKKKKKKOkkOkkkkkkkkkkkkkkkkkkkkkkOOO0000OkO0KKKXXXXXKKKKKKK00k:.
     .xKKKKKKKKXXXKKKKKKKKKKKKKKKOkOOOOkkkkkkkOO000000OkkO00kk0KKXKOkk0KKKXXXXXXXKXXXKK00x,
     .xKKKKKKKXXXXXKKXXKKKXXXKKXKOkkOOOOOkOO000K0OO00OOxdx00xdOKXXXXKOOO0KXXXXXXXXXXXXKK00kc.
     .xKKK000KXXXXXXKKKKKKXXXKKXKOkO0KKKKKKK0O0K0kkkkkOOxdO0kdx0KKKKK0OOO0KXXXXXXXXXXXXKKK00xc,.
     .dK00O0KKXXXXXXXKKKXXXXXXXXKOkOO0KXXXXK0O0KK0O0000OxdxkxddxkkkkkOOOOO00KXXXXXXXXXXXXKKKK00k;
      l0OOO00KKXXXXXXXXXXXXXXXXXKOOO00KK00OOkkkkkkkkkxxddoooooddddxxkkkOOOOO0KKXXXXXXXXXXXXXKKKKOc.
      :000000KXXXXXXXXXXXXXXXXXXK0OO00OkxxdddddddddddddddooddoddddxxkkkkOOOOOO0XXXXXXXXXXXXXXXKKKOc
`
