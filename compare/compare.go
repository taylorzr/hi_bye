package compare

import(
	"github.com/fatih/set"
	"github.com/taylorzr/hibye/root"
)


func FindFired(old []root.User, new []root.User) []root.User {
	return findDifference(old, new)
}

func FindHired(old []root.User, new []root.User) []root.User {
	return findDifference(new, old)
}

func findDifference(from []root.User, to []root.User) []root.User {
	fromSet := idSet(from)
	toSet   := idSet(to)

	differentIDs := set.IntSlice(set.Difference(fromSet, toSet))

	return findUsers(differentIDs, from)
}

func findUsers(ids []int, users []root.User) []root.User {
	foundUsers := []root.User{}

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

func idSet(users []root.User) *set.Set {
	idSet := set.New()

	for _, user := range users {
		idSet.Add(user.ID)
	}

	return idSet
}
