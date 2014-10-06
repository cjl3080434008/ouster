package data

import (
    "testing"
)

func TestReadSMP(t *testing.T) {
    files := []string{
        "adam_c.smp",
        "adam_e.smp",
        "adam_new_c.smp",
        "adam_new_c_n.smp",
        "adam_new_c_s.smp",
        "adam_new_e.smp",
        "adam_new_w.smp",
        "adam_w.smp",
        "alter_of_blood.smp",
        "asylion_dungeon.smp",
        "bathory_battlezone.smp",
        "bathory_dungeon_b1f.smp",
        "bathory_dungeon_b2f.smp",
        "bathory_dungeon_b3f.smp",
        "bathory_dungeon_b4f.smp",
        "bathory_lair.smp",
        "bathory_lair_clon.smp",
        "caligo_dungeon.smp",
        "castalo_ne.smp",
        "castellum.smp",
        "castle_hexserius.smp",
        "castle_octavus.smp",
        "castle_pentanus.smp",
        "castle_quartus.smp",
        "castle_rasen_1_1.smp",
        "castle_rasen_1_2.smp",
        "castle_rasen_2_2.smp",
        "castle_septimus.smp",
        "castle_tertius.smp",
        "clan_hdqrs.smp",
        "devt.smp",
        "drobeta_dungeon_s1f.smp",
        "drobeta_dungeon_v1f.smp",
        "drobeta_ne.smp",
        "drobeta_nw.smp",
        "drobeta_ox.smp",
        "drobeta_se.smp",
        "drobeta_stadium.smp",
        "drobeta_sw.smp",
        "eslania_dungeon.smp",
        "eslania_ne.smp",
        "eslania_nw.smp",
        "eslania_se.smp",
        "eslania_sw.smp",
        "freepk.smp",
        "gate_of_alter.smp",
        "gdr_illusion_01.smp",
        "gdr_illusion_02.smp",
        "gdr_lair_01.smp",
        "gdr_lair_hard.smp",
        "guild_army_1f.smp",
        "guild_army_2f.smp",
        "guild_army_3f.smp",
        "guild_army_4f.smp",
        "guild_army_b1.smp",
        "guild_cleric_1f.smp",
        "guild_cleric_2f.smp",
        "guild_cleric_3f.smp",
        "guild_cleric_4f.smp",
        "guild_cleric_b1.smp",
        "guild_knight_1f.smp",
        "guild_knight_2f.smp",
        "guild_knight_3f.smp",
        "guild_knight_4f.smp",
        "guild_knight_b1.smp",
        "hexserius_dungeon1f.smp",
        "hexserius_dungeon2f.smp",
        "icen_dungeon3f.smp",
        "ik_lab.smp",
        "ik_lab_b1f.smp",
        "ik_lab_b2f.smp",
        "ik_offic.smp",
        "kali_cave.smp",
        "laom_dungeon3f.smp",
        "laom_dungeon4f.smp",
        "laom_dungeon5f.smp",
        "limbo_dungeon.smp",
        "limbo_lair_ne.smp",
        "limbo_lair_nw.smp",
        "limbo_lair_se.smp",
        "limbo_lair_sw.smp",
        "losttaiyan_b1f.smp",
        "losttaiyan_b2f.smp",
        "losttower_1f.smp",
        "losttower_2f.smp",
        "lusttower_1f.smp",
        "lusttower_2f.smp",
        "maze.smp",
        "octavus_dungeon1f.smp",
        "octavus_dungeon2f.smp",
        "ousters_dungeon01.smp",
        "ousters_dungeon02.smp",
        "ousters_dungeon03.smp",
        "ousters_dungeon04.smp",
        "ousters_village.smp",
        "path_to_fears.smp",
        "pentanus_dungeon1f.smp",
        "pentanus_dungeon2f.smp",
        "perona_ne.smp",
        "perona_nw.smp",
        "perona_se.smp",
        "perona_sw.smp",
        "quartus_dungeon1f.smp",
        "quartus_dungeon2f.smp",
        "rasen_battlezone.smp",
        "rasen_training.smp",
        "rasen_yard.smp",
        "rodin_ne.smp",
        "rodin_nw.smp",
        "rodin_se.smp",
        "rodin_sw.smp",
        "septimus_dungeon1f.smp",
        "septimus_dungeon2f.smp",
        "siege_warfare.smp",
        "slayer_battlezone1.smp",
        "slayer_battlezone2.smp",
        "slayerpk.smp",
        "slayers_training.smp",
        "survival.smp",
        "team_hdqrs.smp",
        "tepes_lair.smp",
        "tepes_lair_clon.smp",
        "tertius_dungeon1f.smp",
        "tertius_dungeon2f.smp",
        "timore_ne.smp",
        "timore_nw.smp",
        "timore_se.smp",
        "timore_sw.smp",
        "trapzone01.smp",
        "trapzone02.smp",
        "tunnel_ghorgova.smp",
        "tunnel_peiac.smp",
        "tutorial_n.smp",
        "tutorial_s.smp",
        "under_pass_1f.smp",
        "under_pass_2f.smp",
        "vampire_village.smp",
        "vampirepk.smp",
        "vranco_ne.smp",
        "vranco_nw.smp",
        "vranco_se.smp",
        "vranco_sw.smp",
    }

    for _, v := range files {
        smp, err := ReadSMP(v)
        if err != nil {
            t.Fatal(err)
        }
        if smp == nil {
            t.Fatalf("err不为空，smp却为空:%s\n", v)
        }
        // t.Logf("name: %s, id: %d\n", v, smp.ZoneID)
        // t.Logf("%#v\n", smp)
        if smp.Width == 0 || smp.Height == 0 {
            t.Fatal("这个地图不对：%s\n", v)
        }
    }
}

func TestReadSSI(t *testing.T) {
    files := []string{
        "adam_c.ssi",
        "adam_e.ssi",
        "adam_new_c.ssi",
        "adam_new_c_n.ssi",
        "adam_new_c_s.ssi",
        "adam_new_e.ssi",
        "adam_new_w.ssi",
        "adam_w.ssi",
        "alter_of_blood.ssi",
        "asylion_dungeon.ssi",
        "bathory_battlezone.ssi",
        "bathory_dungeon_b1f.ssi",
        "bathory_dungeon_b2f.ssi",
        "bathory_dungeon_b3f.ssi",
        "bathory_dungeon_b4f.ssi",
        "bathory_lair.ssi",
        "bathory_lair_clon.ssi",
        "caligo_dungeon.ssi",
        "castalo_ne.ssi",
        "castellum.ssi",
        "castle_hexserius.ssi",
        "castle_octavus.ssi",
        "castle_pentanus.ssi",
        "castle_quartus.ssi",
        "castle_rasen_1_1.ssi",
        "castle_rasen_1_2.ssi",
        "castle_rasen_2_2.ssi",
        "castle_septimus.ssi",
        "castle_tertius.ssi",
        "clan_hdqrs.ssi",
        "devt.ssi",
        "drobeta_dungeon_s1f.ssi",
        "drobeta_dungeon_v1f.ssi",
        "drobeta_ne.ssi",
        "drobeta_nw.ssi",
        "drobeta_ox.ssi",
        "drobeta_se.ssi",
        "drobeta_stadium.ssi",
        "drobeta_sw.ssi",
        "eslania_dungeon.ssi",
        "eslania_ne.ssi",
        "eslania_nw.ssi",
        "eslania_se.ssi",
        "eslania_sw.ssi",
        "freepk.ssi",
        "gate_of_alter.ssi",
        "gdr_illusion_01.ssi",
        "gdr_illusion_02.ssi",
        "gdr_lair_01.ssi",
        "gdr_lair_hard.ssi",
        "guild_army_1f.ssi",
        "guild_army_2f.ssi",
        "guild_army_3f.ssi",
        "guild_army_4f.ssi",
        "guild_army_b1.ssi",
        "guild_cleric_1f.ssi",
        "guild_cleric_2f.ssi",
        "guild_cleric_3f.ssi",
        "guild_cleric_4f.ssi",
        "guild_cleric_b1.ssi",
        "guild_knight_1f.ssi",
        "guild_knight_2f.ssi",
        "guild_knight_3f.ssi",
        "guild_knight_4f.ssi",
        "guild_knight_b1.ssi",
        "hexserius_dungeon1f.ssi",
        "hexserius_dungeon2f.ssi",
        "icen_dungeon3f.ssi",
        "ik_lab.ssi",
        "ik_lab_b1f.ssi",
        "ik_lab_b2f.ssi",
        "ik_offic.ssi",
        "kali_cave.ssi",
        "laom_dungeon3f.ssi",
        "laom_dungeon4f.ssi",
        "laom_dungeon5f.ssi",
        "limbo_dungeon.ssi",
        "limbo_lair_ne.ssi",
        "limbo_lair_nw.ssi",
        "limbo_lair_se.ssi",
        "limbo_lair_sw.ssi",
        "losttaiyan_b1f.ssi",
        "losttaiyan_b2f.ssi",
        "losttower_1f.ssi",
        "losttower_2f.ssi",
        "lusttower_1f.ssi",
        "lusttower_2f.ssi",
        "maze.ssi",
        "octavus_dungeon1f.ssi",
        "octavus_dungeon2f.ssi",
        "ousters_dungeon01.ssi",
        "ousters_dungeon02.ssi",
        "ousters_dungeon03.ssi",
        "ousters_dungeon04.ssi",
        "ousters_village.ssi",
        "path_to_fears.ssi",
        "pentanus_dungeon1f.ssi",
        "pentanus_dungeon2f.ssi",
        "perona_ne.ssi",
        "perona_nw.ssi",
        "perona_se.ssi",
        "perona_sw.ssi",
        "quartus_dungeon1f.ssi",
        "quartus_dungeon2f.ssi",
        "rasen_battlezone.ssi",
        "rasen_training.ssi",
        "rasen_yard.ssi",
        "rodin_ne.ssi",
        "rodin_nw.ssi",
        "rodin_se.ssi",
        "rodin_sw.ssi",
        "septimus_dungeon1f.ssi",
        "septimus_dungeon2f.ssi",
        "siege_warfare.ssi",
        "slayer_battlezone1.ssi",
        "slayer_battlezone2.ssi",
        "slayerpk.ssi",
        "slayers_training.ssi",
        "survival.ssi",
        "team_hdqrs.ssi",
        "tepes_lair.ssi",
        "tepes_lair_clon.ssi",
        "tertius_dungeon1f.ssi",
        "tertius_dungeon2f.ssi",
        "timore_ne.ssi",
        "timore_nw.ssi",
        "timore_se.ssi",
        "timore_sw.ssi",
        "trapzone01.ssi",
        "trapzone02.ssi",
        "tunnel_ghorgova.ssi",
        "tunnel_peiac.ssi",
        "tutorial_n.ssi",
        "tutorial_s.ssi",
        "under_pass_1f.ssi",
        "under_pass_2f.ssi",
        "vampire_village.ssi",
        "vampirepk.ssi",
        "vranco_ne.ssi",
        "vranco_nw.ssi",
        "vranco_se.ssi",
        "vranco_sw.ssi",
    }

    for _, v := range files {
        _, err := ReadSSI(v)
        if err != nil {
            t.Fatal(err)
        }
    }
}
