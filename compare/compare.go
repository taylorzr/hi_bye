package compare

import(
	"github.com/fatih/set"
	"github.com/taylorzr/hibye/root"
)


func FindFired(old root.Users, new root.Users) root.Users {
	return findDifference(old, new)
}

func FindHired(old root.Users, new root.Users) root.Users {
	return findDifference(new, old)
}

func findDifference(from root.Users, to root.Users) root.Users {
	fromSet := idSet(from)
	toSet   := idSet(to)

	differentIDs := set.IntSlice(set.Difference(fromSet, toSet))

	return findUsers(differentIDs, from)
}

func findUsers(ids []int, users root.Users) root.Users {
	foundUsers := root.Users{}

	if len(ids) > 0 {
		usersByID := map[int]root.User{}

		for _, user := range users {
			usersByID[user.ID] = user
		}

		for _, id := range ids {
			foundUsers = append(foundUsers, usersByID[id])
		}
	}

	return foundUsers
}

func idSet(users root.Users) *set.Set {
	idSet := set.New()

	for _, user := range users {
		idSet.Add(user.ID)
	}

	return idSet
}
