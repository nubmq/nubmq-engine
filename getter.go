package main

import (
	"fmt"
)

func getAtIndex(idx int, key string, keeper *ShardManagerKeeperTemp) (string, bool) {
	SMidx, localIdx := getShardNumberAndIndexPair(idx)
	keeper.ShardManagers[SMidx].mutex.RLock()
	if SMidx >= len(keeper.ShardManagers) {
		keeper.ShardManagers[SMidx].mutex.RUnlock()
		fmt.Println("Looking too far, not found")
		return "NaN", false
	}
	targetSM := keeper.ShardManagers[SMidx]
	keeper.ShardManagers[SMidx].mutex.RUnlock()

	targetSM.mutex.RLock()
	target := targetSM.Shards[localIdx]
	targetSM.mutex.RUnlock()
	value, ok := target.data.Load(key)
	if ok {
		entry := (value.(Entry)).value
		return entry, true
	} else {
		fmt.Println("just not there man", key)
		return "NaN", false
	}
}

func _getKey(key string) (string, bool) {
	// check the new table first then the old one
	// TODO: can we do something better ? having to check two tables to fulfil one request is slow

	if true {
		newShardManagerKeeper.mutex.RLock()

		ret, found := getAtIndex(getKeyHash(key, &newShardManagerKeeper), key, &newShardManagerKeeper)

		newShardManagerKeeper.mutex.RUnlock()

		if found {
			return ret, found
		}
	}

	ShardManagerKeeper.mutex.RLock()

	ret, found := getAtIndex(getKeyHash(key, &ShardManagerKeeper), key, &ShardManagerKeeper)

	ShardManagerKeeper.mutex.RUnlock()

	return ret, found
}
