package blogposts_test

import (
	blogposts "go-with-tests/reading-files"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody  = "Title: Post 1\nDescription: Description 1\nTags: tdd, go\n---\nHello\nWorld"
		secondBody = "Title: Post 2\nDescription: Description 2\nTags: rust, borrow-checker\n---\nB\nL\nM"
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	assertNoError(t, err)

	assertPostsLength(t, posts, fs)

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body:        "Hello\nWorld",
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func assertPostsLength(t *testing.T, posts []blogposts.Post, fs fstest.MapFS) {
	t.Helper()
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
