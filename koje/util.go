package koje

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dgraph-io/badger"
)

func _t(en string, tp string, m *discordgo.MessageCreate) string {
	if chanTped(m.ChannelID) {
		return tp
	} else {
		return en
	}
}

func chanTped(id string) bool {
	var ret bool
	db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id + "_tp"))
		if err != nil {
			ret = false
			return err
		}
		item.Value(func(val []byte) error {
			if string(val) == "true" {
				ret = true
			} else {
				ret = false
			}
			return nil
		})
		return nil
	})
	return ret
}

func setChanTped(id string, blacklisted bool) {
	db.Update(func(txn *badger.Txn) error {
		if blacklisted {
			err := txn.Set([]byte(id+"_tp"), []byte("true"))
			return err
		} else {
			err := txn.Set([]byte(id+"_tp"), []byte("false"))
			return err
		}
	})
}

func chanBlacklisted(id string) bool {
	var ret bool
	db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id + "_blacklist"))
		if err != nil {
			ret = false
			return err
		}
		item.Value(func(val []byte) error {
			if string(val) == "true" {
				ret = true
			} else {
				ret = false
			}
			return nil
		})
		return nil
	})
	return ret
}

func setChanBlacklisted(id string, blacklisted bool) {
	db.Update(func(txn *badger.Txn) error {
		if blacklisted {
			err := txn.Set([]byte(id+"_blacklist"), []byte("true"))
			return err
		} else {
			err := txn.Set([]byte(id+"_blacklist"), []byte("false"))
			return err
		}
	})
}
