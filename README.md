# Render open source fonts to polygons in Go

This is an experimental package used to render open source fonts to
polygons using Go.

## Example usage

To use one or more fonts within a Go program, import the main
package and the font(s) you want, like this:

```go
import (
  "github.com/gmlewis/go-fonts/fonts"
  _ "github.com/gmlewis/go-fonts/fonts/ubuntumonoregular"
)
```

Then render the text to polygons and use them however you want:

```go
  render, err := fonts.Text(xPos, yPos, xScale, yScale, message, "ubuntumonoregular")
  if err != nil {
    return err
  }
  log.Printf("MBB: (%.2f,%.2f)-(%.2f,%.2f)", render.Xmin, render.Ymin,render.Xmax, render.Ymax)
  for _, poly := range render.Polygons {
    // ...
  }
```

See https://github.com/gmlewis/go-gerber for an example application
that uses this package.

## Status
[![GoDoc](https://godoc.org/github.com/gmlewis/go-fonts/fonts?status.svg)](https://godoc.org/github.com/gmlewis/go-fonts/fonts)
[![Build Status](https://travis-ci.org/gmlewis/go-fonts.png)](https://travis-ci.org/gmlewis/go-fonts)

----------------------------------------------------------------------

Enjoy!

----------------------------------------------------------------------

# License

Copyright 2019 Glenn M. Lewis. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

----------------------------------------------------------------------

## Font samples
[![aaarghnormal](images/sample_aaarghnormal.png)](fonts/aaarghnormal)
[![abeezee_italic](images/sample_abeezee_italic.png)](fonts/abeezee_italic)
[![abeezee_regular](images/sample_abeezee_regular.png)](fonts/abeezee_regular)
[![acme_regular](images/sample_acme_regular.png)](fonts/acme_regular)
[![aguafinascript_regular](images/sample_aguafinascript_regular.png)](fonts/aguafinascript_regular)
[![aileron_black](images/sample_aileron_black.png)](fonts/aileron_black)
[![aileron_blackitalic](images/sample_aileron_blackitalic.png)](fonts/aileron_blackitalic)
[![aileron_bold](images/sample_aileron_bold.png)](fonts/aileron_bold)
[![aileron_bolditalic](images/sample_aileron_bolditalic.png)](fonts/aileron_bolditalic)
[![aileron_heavy](images/sample_aileron_heavy.png)](fonts/aileron_heavy)
[![aileron_heavyitalic](images/sample_aileron_heavyitalic.png)](fonts/aileron_heavyitalic)
[![aileron_italic](images/sample_aileron_italic.png)](fonts/aileron_italic)
[![aileron_light](images/sample_aileron_light.png)](fonts/aileron_light)
[![aileron_lightitalic](images/sample_aileron_lightitalic.png)](fonts/aileron_lightitalic)
[![aileron_regular](images/sample_aileron_regular.png)](fonts/aileron_regular)
[![aileron_semibold](images/sample_aileron_semibold.png)](fonts/aileron_semibold)
[![aileron_semibolditalic](images/sample_aileron_semibolditalic.png)](fonts/aileron_semibolditalic)
[![aileron_thin](images/sample_aileron_thin.png)](fonts/aileron_thin)
[![aileron_thinitalic](images/sample_aileron_thinitalic.png)](fonts/aileron_thinitalic)
[![aileron_ultralight](images/sample_aileron_ultralight.png)](fonts/aileron_ultralight)
[![aileron_ultralightitalic](images/sample_aileron_ultralightitalic.png)](fonts/aileron_ultralightitalic)
[![airstream](images/sample_airstream.png)](fonts/airstream)
[![alexbrush_regular](images/sample_alexbrush_regular.png)](fonts/alexbrush_regular)
[![allura_regular](images/sample_allura_regular.png)](fonts/allura_regular)
[![amerika](images/sample_amerika.png)](fonts/amerika)
[![amerikasans](images/sample_amerikasans.png)](fonts/amerikasans)
[![amita_bold](images/sample_amita_bold.png)](fonts/amita_bold)
[![amita_regular](images/sample_amita_regular.png)](fonts/amita_regular)
[![anagram](images/sample_anagram.png)](fonts/anagram)
[![aquilinetwo](images/sample_aquilinetwo.png)](fonts/aquilinetwo)
[![arizonia_regular](images/sample_arizonia_regular.png)](fonts/arizonia_regular)
[![asset_regular](images/sample_asset_regular.png)](fonts/asset_regular)
[![atomicage_regular](images/sample_atomicage_regular.png)](fonts/atomicage_regular)
[![baloo](images/sample_baloo.png)](fonts/baloo)
[![baskervville_italic](images/sample_baskervville_italic.png)](fonts/baskervville_italic)
[![baskervville_regular](images/sample_baskervville_regular.png)](fonts/baskervville_regular)
[![berkshireswash_regular](images/sample_berkshireswash_regular.png)](fonts/berkshireswash_regular)
[![bitstreamverasansmono_bold](images/sample_bitstreamverasansmono_bold.png)](fonts/bitstreamverasansmono_bold)
[![bitstreamverasansmono_boldob](images/sample_bitstreamverasansmono_boldob.png)](fonts/bitstreamverasansmono_boldob)
[![bitstreamverasansmono_oblique](images/sample_bitstreamverasansmono_oblique.png)](fonts/bitstreamverasansmono_oblique)
[![bitstreamverasansmono_roman](images/sample_bitstreamverasansmono_roman.png)](fonts/bitstreamverasansmono_roman)
[![blazium](images/sample_blazium.png)](fonts/blazium)
[![cacchampagne](images/sample_cacchampagne.png)](fonts/cacchampagne)
[![carrelectronicdingbats](images/sample_carrelectronicdingbats.png)](fonts/carrelectronicdingbats)
[![coiny_regular](images/sample_coiny_regular.png)](fonts/coiny_regular)
[![combinumeralsltd](images/sample_combinumeralsltd.png)](fonts/combinumeralsltd)
[![davysregular](images/sample_davysregular.png)](fonts/davysregular)
[![deliusswashcaps_regular](images/sample_deliusswashcaps_regular.png)](fonts/deliusswashcaps_regular)
[![diplomata_regular](images/sample_diplomata_regular.png)](fonts/diplomata_regular)
[![dited](images/sample_dited.png)](fonts/dited)
[![embossedblack_normal](images/sample_embossedblack_normal.png)](fonts/embossedblack_normal)
[![embossedblackwide_normal](images/sample_embossedblackwide_normal.png)](fonts/embossedblackwide_normal)
[![englandhanddb](images/sample_englandhanddb.png)](fonts/englandhanddb)
[![entypo](images/sample_entypo.png)](fonts/entypo)
[![f20db](images/sample_f20db.png)](fonts/f20db)
[![f2dumb](images/sample_f2dumb.png)](fonts/f2dumb)
[![f3dumb](images/sample_f3dumb.png)](fonts/f3dumb)
[![fantasquesansmono_bold](images/sample_fantasquesansmono_bold.png)](fonts/fantasquesansmono_bold)
[![fantasquesansmono_bolditalic](images/sample_fantasquesansmono_bolditalic.png)](fonts/fantasquesansmono_bolditalic)
[![fantasquesansmono_italic](images/sample_fantasquesansmono_italic.png)](fonts/fantasquesansmono_italic)
[![fantasquesansmono_regular](images/sample_fantasquesansmono_regular.png)](fonts/fantasquesansmono_regular)
[![fascinate_inlineregular](images/sample_fascinate_inlineregular.png)](fonts/fascinate_inlineregular)
[![fauxsnowbrk](images/sample_fauxsnowbrk.png)](fonts/fauxsnowbrk)
[![floralia](images/sample_floralia.png)](fonts/floralia)
[![font3933](images/sample_font3933.png)](fonts/font3933)
[![fontleroybrown](images/sample_fontleroybrown.png)](fonts/fontleroybrown)
[![freebooterscript](images/sample_freebooterscript.png)](fonts/freebooterscript)
[![freemono](images/sample_freemono.png)](fonts/freemono)
[![freemonobold](images/sample_freemonobold.png)](fonts/freemonobold)
[![freemonoboldoblique](images/sample_freemonoboldoblique.png)](fonts/freemonoboldoblique)
[![freemonooblique](images/sample_freemonooblique.png)](fonts/freemonooblique)
[![freesans](images/sample_freesans.png)](fonts/freesans)
[![freesansbold](images/sample_freesansbold.png)](fonts/freesansbold)
[![freesansboldoblique](images/sample_freesansboldoblique.png)](fonts/freesansboldoblique)
[![freesansoblique](images/sample_freesansoblique.png)](fonts/freesansoblique)
[![freeserif](images/sample_freeserif.png)](fonts/freeserif)
[![freeserifbold](images/sample_freeserifbold.png)](fonts/freeserifbold)
[![freeserifbolditalic](images/sample_freeserifbolditalic.png)](fonts/freeserifbolditalic)
[![freeserifitalic](images/sample_freeserifitalic.png)](fonts/freeserifitalic)
[![genzschetheyse](images/sample_genzschetheyse.png)](fonts/genzschetheyse)
[![genzschetheysealternate](images/sample_genzschetheysealternate.png)](fonts/genzschetheysealternate)
[![geometrysoftpro_boldn](images/sample_geometrysoftpro_boldn.png)](fonts/geometrysoftpro_boldn)
[![gooddogregular](images/sample_gooddogregular.png)](fonts/gooddogregular)
[![goudystm](images/sample_goudystm.png)](fonts/goudystm)
[![goudystm_italic](images/sample_goudystm_italic.png)](fonts/goudystm_italic)
[![grandhotel_regular](images/sample_grandhotel_regular.png)](fonts/grandhotel_regular)
[![greatvibes_regular](images/sample_greatvibes_regular.png)](fonts/greatvibes_regular)
[![grutchshaded](images/sample_grutchshaded.png)](fonts/grutchshaded)
[![hanaleifill_regular](images/sample_hanaleifill_regular.png)](fonts/hanaleifill_regular)
[![hanalei_regular](images/sample_hanalei_regular.png)](fonts/hanalei_regular)
[![headhunter_regular](images/sample_headhunter_regular.png)](fonts/headhunter_regular)
[![heavydata](images/sample_heavydata.png)](fonts/heavydata)
[![helsinkiregular](images/sample_helsinkiregular.png)](fonts/helsinkiregular)
[![heydingsicons](images/sample_heydingsicons.png)](fonts/heydingsicons)
[![im_fell_flowers_1](images/sample_im_fell_flowers_1.png)](fonts/im_fell_flowers_1)
[![im_fell_flowers_2](images/sample_im_fell_flowers_2.png)](fonts/im_fell_flowers_2)
[![impactlabel](images/sample_impactlabel.png)](fonts/impactlabel)
[![impactlabelreversed](images/sample_impactlabelreversed.png)](fonts/impactlabelreversed)
[![incisedblack_normal](images/sample_incisedblack_normal.png)](fonts/incisedblack_normal)
[![incisedblackwide_normal](images/sample_incisedblackwide_normal.png)](fonts/incisedblackwide_normal)
[![inconsolata](images/sample_inconsolata.png)](fonts/inconsolata)
[![italianno_regular](images/sample_italianno_regular.png)](fonts/italianno_regular)
[![kawoszeh](images/sample_kawoszeh.png)](fonts/kawoszeh)
[![kellssd](images/sample_kellssd.png)](fonts/kellssd)
[![kingthingsitalique](images/sample_kingthingsitalique.png)](fonts/kingthingsitalique)
[![kingthingsxstitch](images/sample_kingthingsxstitch.png)](fonts/kingthingsxstitch)
[![konstytucyja](images/sample_konstytucyja.png)](fonts/konstytucyja)
[![landliebe](images/sample_landliebe.png)](fonts/landliebe)
[![latoregular](images/sample_latoregular.png)](fonts/latoregular)
[![leaguescriptthin_regular](images/sample_leaguescriptthin_regular.png)](fonts/leaguescriptthin_regular)
[![lobstertwo](images/sample_lobstertwo.png)](fonts/lobstertwo)
[![lobstertwo_bold](images/sample_lobstertwo_bold.png)](fonts/lobstertwo_bold)
[![lobstertwo_bolditalic](images/sample_lobstertwo_bolditalic.png)](fonts/lobstertwo_bolditalic)
[![lobstertwo_italic](images/sample_lobstertwo_italic.png)](fonts/lobstertwo_italic)
[![loversquarrel_regular](images/sample_loversquarrel_regular.png)](fonts/loversquarrel_regular)
[![membra](images/sample_membra.png)](fonts/membra)
[![miama](images/sample_miama.png)](fonts/miama)
[![modak](images/sample_modak.png)](fonts/modak)
[![monospacetypewriter](images/sample_monospacetypewriter.png)](fonts/monospacetypewriter)
[![monoton_regular](images/sample_monoton_regular.png)](fonts/monoton_regular)
[![montez_regular](images/sample_montez_regular.png)](fonts/montez_regular)
[![mothproofscript](images/sample_mothproofscript.png)](fonts/mothproofscript)
[![mutluornamental](images/sample_mutluornamental.png)](fonts/mutluornamental)
[![overlockregular](images/sample_overlockregular.png)](fonts/overlockregular)
[![oxygen](images/sample_oxygen.png)](fonts/oxygen)
[![oxygen_bold](images/sample_oxygen_bold.png)](fonts/oxygen_bold)
[![oxygen_bolditalic](images/sample_oxygen_bolditalic.png)](fonts/oxygen_bolditalic)
[![oxygen_italic](images/sample_oxygen_italic.png)](fonts/oxygen_italic)
[![oxygenmono_regular](images/sample_oxygenmono_regular.png)](fonts/oxygenmono_regular)
[![pacifico](images/sample_pacifico.png)](fonts/pacifico)
[![parisienne_regular](images/sample_parisienne_regular.png)](fonts/parisienne_regular)
[![plainblack_normal](images/sample_plainblack_normal.png)](fonts/plainblack_normal)
[![plainblackwide_normal](images/sample_plainblackwide_normal.png)](fonts/plainblackwide_normal)
[![plasmadripbrk](images/sample_plasmadripbrk.png)](fonts/plasmadripbrk)
[![plasmadripemptybrk](images/sample_plasmadripemptybrk.png)](fonts/plasmadripemptybrk)
[![plexifontbv](images/sample_plexifontbv.png)](fonts/plexifontbv)
[![princesssofia](images/sample_princesssofia.png)](fonts/princesssofia)
[![printbold](images/sample_printbold.png)](fonts/printbold)
[![printclearly](images/sample_printclearly.png)](fonts/printclearly)
[![printdashed](images/sample_printdashed.png)](fonts/printdashed)
[![printersornamentsone](images/sample_printersornamentsone.png)](fonts/printersornamentsone)
[![promocyja](images/sample_promocyja.png)](fonts/promocyja)
[![qwigley_regular](images/sample_qwigley_regular.png)](fonts/qwigley_regular)
[![ralewaydots_regular](images/sample_ralewaydots_regular.png)](fonts/ralewaydots_regular)
[![rechtmanplain](images/sample_rechtmanplain.png)](fonts/rechtmanplain)
[![rothenburgdecorative_normal](images/sample_rothenburgdecorative_normal.png)](fonts/rothenburgdecorative_normal)
[![rougescript_regular](images/sample_rougescript_regular.png)](fonts/rougescript_regular)
[![rubik_black](images/sample_rubik_black.png)](fonts/rubik_black)
[![rubik_blackitalic](images/sample_rubik_blackitalic.png)](fonts/rubik_blackitalic)
[![rubik_bold](images/sample_rubik_bold.png)](fonts/rubik_bold)
[![rubik_bolditalic](images/sample_rubik_bolditalic.png)](fonts/rubik_bolditalic)
[![rubik_italic](images/sample_rubik_italic.png)](fonts/rubik_italic)
[![rubik_light](images/sample_rubik_light.png)](fonts/rubik_light)
[![rubik_lightitalic](images/sample_rubik_lightitalic.png)](fonts/rubik_lightitalic)
[![rubik_medium](images/sample_rubik_medium.png)](fonts/rubik_medium)
[![rubik_mediumitalic](images/sample_rubik_mediumitalic.png)](fonts/rubik_mediumitalic)
[![rubik_regular](images/sample_rubik_regular.png)](fonts/rubik_regular)
[![ruthie_regular](images/sample_ruthie_regular.png)](fonts/ruthie_regular)
[![rye_regular](images/sample_rye_regular.png)](fonts/rye_regular)
[![sail_regular](images/sample_sail_regular.png)](fonts/sail_regular)
[![satisfy_regular](images/sample_satisfy_regular.png)](fonts/satisfy_regular)
[![scratch](images/sample_scratch.png)](fonts/scratch)
[![scriptinapro](images/sample_scriptinapro.png)](fonts/scriptinapro)
[![sevillana_regular](images/sample_sevillana_regular.png)](fonts/sevillana_regular)
[![sfarcheryblack](images/sample_sfarcheryblack.png)](fonts/sfarcheryblack)
[![sfarcheryblack_oblique](images/sample_sfarcheryblack_oblique.png)](fonts/sfarcheryblack_oblique)
[![sfarcheryblacksc](images/sample_sfarcheryblacksc.png)](fonts/sfarcheryblacksc)
[![sfarcheryblacksc_oblique](images/sample_sfarcheryblacksc_oblique.png)](fonts/sfarcheryblacksc_oblique)
[![sfwasabi](images/sample_sfwasabi.png)](fonts/sfwasabi)
[![sfwasabicondensed](images/sample_sfwasabicondensed.png)](fonts/sfwasabicondensed)
[![shojumaru_regular](images/sample_shojumaru_regular.png)](fonts/shojumaru_regular)
[![shortstack](images/sample_shortstack.png)](fonts/shortstack)
[![simpel_medium](images/sample_simpel_medium.png)](fonts/simpel_medium)
[![sjonarbok_classic](images/sample_sjonarbok_classic.png)](fonts/sjonarbok_classic)
[![snickles](images/sample_snickles.png)](fonts/snickles)
[![sniglet_extrabold](images/sample_sniglet_extrabold.png)](fonts/sniglet_extrabold)
[![sniglet_regular](images/sample_sniglet_regular.png)](fonts/sniglet_regular)
[![sofia_regular](images/sample_sofia_regular.png)](fonts/sofia_regular)
[![solveigbold](images/sample_solveigbold.png)](fonts/solveigbold)
[![solveigbold_italic](images/sample_solveigbold_italic.png)](fonts/solveigbold_italic)
[![solveigdemibold](images/sample_solveigdemibold.png)](fonts/solveigdemibold)
[![solveigdemibold_italic](images/sample_solveigdemibold_italic.png)](fonts/solveigdemibold_italic)
[![solveigdisplay](images/sample_solveigdisplay.png)](fonts/solveigdisplay)
[![solveigdisplay_italic](images/sample_solveigdisplay_italic.png)](fonts/solveigdisplay_italic)
[![solveigtext](images/sample_solveigtext.png)](fonts/solveigtext)
[![solveigtext_italic](images/sample_solveigtext_italic.png)](fonts/solveigtext_italic)
[![sonsieone](images/sample_sonsieone.png)](fonts/sonsieone)
[![soria_soria](images/sample_soria_soria.png)](fonts/soria_soria)
[![soucisans](images/sample_soucisans.png)](fonts/soucisans)
[![spacemono_bold](images/sample_spacemono_bold.png)](fonts/spacemono_bold)
[![spacemono_bolditalic](images/sample_spacemono_bolditalic.png)](fonts/spacemono_bolditalic)
[![spacemono_italic](images/sample_spacemono_italic.png)](fonts/spacemono_italic)
[![spacemono_regular](images/sample_spacemono_regular.png)](fonts/spacemono_regular)
[![spiltink](images/sample_spiltink.png)](fonts/spiltink)
[![spirax_regular](images/sample_spirax_regular.png)](fonts/spirax_regular)
[![sportrop_regular](images/sample_sportrop_regular.png)](fonts/sportrop_regular)
[![squadaone_regular](images/sample_squadaone_regular.png)](fonts/squadaone_regular)
[![stardosstencil_bold](images/sample_stardosstencil_bold.png)](fonts/stardosstencil_bold)
[![stardosstencil_regular](images/sample_stardosstencil_regular.png)](fonts/stardosstencil_regular)
[![stateface_regular](images/sample_stateface_regular.png)](fonts/stateface_regular)
[![stmarie_thin](images/sample_stmarie_thin.png)](fonts/stmarie_thin)
[![symbolsigns_basisset](images/sample_symbolsigns_basisset.png)](fonts/symbolsigns_basisset)
[![synthetiqueot](images/sample_synthetiqueot.png)](fonts/synthetiqueot)
[![tangerine](images/sample_tangerine.png)](fonts/tangerine)
[![tangerine_bold](images/sample_tangerine_bold.png)](fonts/tangerine_bold)
[![teutonicno1_demibold](images/sample_teutonicno1_demibold.png)](fonts/teutonicno1_demibold)
[![teutonicno2_demibold](images/sample_teutonicno2_demibold.png)](fonts/teutonicno2_demibold)
[![teutonicno3_demibold](images/sample_teutonicno3_demibold.png)](fonts/teutonicno3_demibold)
[![teutonicno4_demibold](images/sample_teutonicno4_demibold.png)](fonts/teutonicno4_demibold)
[![texgyreadventor_bold](images/sample_texgyreadventor_bold.png)](fonts/texgyreadventor_bold)
[![texgyreadventor_bolditalic](images/sample_texgyreadventor_bolditalic.png)](fonts/texgyreadventor_bolditalic)
[![texgyreadventor_italic](images/sample_texgyreadventor_italic.png)](fonts/texgyreadventor_italic)
[![texgyreadventor_regular](images/sample_texgyreadventor_regular.png)](fonts/texgyreadventor_regular)
[![topsecret_bold](images/sample_topsecret_bold.png)](fonts/topsecret_bold)
[![typemymusic_notation](images/sample_typemymusic_notation.png)](fonts/typemymusic_notation)
[![ubuntumonoregular](images/sample_ubuntumonoregular.png)](fonts/ubuntumonoregular)
[![vanilla](images/sample_vanilla.png)](fonts/vanilla)
[![vastshadow_regular](images/sample_vastshadow_regular.png)](fonts/vastshadow_regular)
[![veggieburger](images/sample_veggieburger.png)](fonts/veggieburger)
[![veggieburger_bold](images/sample_veggieburger_bold.png)](fonts/veggieburger_bold)
[![veggieburger_light](images/sample_veggieburger_light.png)](fonts/veggieburger_light)
[![veterantypewriter](images/sample_veterantypewriter.png)](fonts/veterantypewriter)
[![vibur](images/sample_vibur.png)](fonts/vibur)
[![wcsoldoutabta](images/sample_wcsoldoutabta.png)](fonts/wcsoldoutabta)
[![wcsoldoutbbta](images/sample_wcsoldoutbbta.png)](fonts/wcsoldoutbbta)
[![wcsoldoutcbta](images/sample_wcsoldoutcbta.png)](fonts/wcsoldoutcbta)
[![websymbols_regular](images/sample_websymbols_regular.png)](fonts/websymbols_regular)
[![wellfleet_regular](images/sample_wellfleet_regular.png)](fonts/wellfleet_regular)
[![windsong](images/sample_windsong.png)](fonts/windsong)
[![woodennickelblack](images/sample_woodennickelblack.png)](fonts/woodennickelblack)
[![yataghan](images/sample_yataghan.png)](fonts/yataghan)
[![yellowtail](images/sample_yellowtail.png)](fonts/yellowtail)
[![yesevaone](images/sample_yesevaone.png)](fonts/yesevaone)
[![zenda](images/sample_zenda.png)](fonts/zenda)
[![znikomit](images/sample_znikomit.png)](fonts/znikomit)
[![znikomitno24](images/sample_znikomitno24.png)](fonts/znikomitno24)
