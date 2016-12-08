// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package tumblr

import (
	"encoding/json"
	"errors"
	"net/url"
	"time"
)

type URL struct {
	url.URL
}

func (u URL) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

func (u *URL) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	uu, err := url.Parse(str)
	if err != nil {
		return err
	}

	u.URL = *uu
	return nil
}

func (u URL) String() string {
	return u.URL.String()
}

type UnixTime struct {
	time.Time
}

func (t UnixTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Unix())
}

func (t *UnixTime) UnmarshalJSON(b []byte) error {
	var unix int64
	if err := json.Unmarshal(b, &unix); err != nil {
		return err
	}

	t.Time = time.Unix(unix, 0)
	return nil
}

type Response struct {
	Meta struct {
		Status int
		Msg    string
	}

	Response json.RawMessage
}

type PostsResponse struct {
	Blog       Blog
	Posts      []json.RawMessage
	TotalPosts int `json:"total_posts"`
}

type Blog struct {
	Title       string
	Posts       int
	Name        string
	Updated     UnixTime
	Description string
	Ask         bool
	AskAnon     bool `json:"ask_anon"`
	Likes       int

	Url    URL
	IsNSFW bool `json:"is_nsfw"`
}

type BasePost struct {
	Blogame   string
	Id        int
	PostUrl   URL `json:"post_url"`
	Type      string
	Timestamp UnixTime
	//Date        time.Time
	Format      string
	ReblogKey   string `json:"reblog_key"`
	Tags        []string
	Bookmarklet bool
	Mobile      bool
	SourceUrl   URL    `json:"source_title"`
	SourceTitle string `json:"source_title"`
	Liked       bool
	State       string
	//TotalPosts  int
}

type TextPost struct {
	BasePost

	Title string
	Body  string
}

type PhotoPost struct {
	BasePost

	Photos  []PostPhoto
	Caption string
	Width   int
	Height  int
}

type PostPhoto struct {
	Caption      string
	OriginalSize PostPhotoSize   `json:"original_size"`
	AltSizes     []PostPhotoSize `json:"alt_sizes"`
}

type PostPhotoSize struct {
	Width  int
	Height int
	Url    URL
}

type QuotePost struct {
	BasePost

	Text   string
	Source string
}

type LinkPost struct {
	BasePost

	Title       string
	Url         URL
	Author      string
	Excerpt     string
	Publisher   string
	Photots     []LinkPhoto
	Description string
}

type LinkPhoto struct {
	Caption      string
	OriginalSize LinkPhotoSize   `json:"original_size"`
	AltSizes     []LinkPhotoSize `json:"alt_sizes"`
}

type LinkPhotoSize struct {
	Width  int
	Height int
	Url    URL
}

type ChatPost struct {
	BasePost

	Title    string
	Body     string
	Dialogue []ChatDialogue
}

type ChatDialogue struct {
	Name   string
	Label  string
	Phrase string
}

type AudioPost struct {
	BasePost

	Caption     string
	Player      string
	Plays       int
	AlbumArt    string `json:"album_art"`
	Artist      string
	Album       string
	TrackName   string `json:"track_name"`
	TrackNumber int    `json:"track_number"`
	Year        int
}

type VideoPost struct {
	BasePost

	Caption string
	Player  []VideoPlayer
}

type VideoPlayer struct {
	Width     int
	EmbedCode string `json:"embed_code"`
}

type AnswerPost struct {
	BasePost

	AskingName string `json:"asking_name"`
	AskingUrl  URL    `json:"asking_url"`
	Question   string
	Answer     string
}

func UnmarshalPost(b []byte, post *interface{}) error {
	var base BasePost
	if err := json.Unmarshal(b, &base); err != nil {
		return err
	}

	switch base.Type {
	case "text":
		*post = new(TextPost)
	case "quote":
		*post = new(QuotePost)
	case "link":
		*post = new(LinkPost)
	case "answer":
		*post = new(AnswerPost)
	case "video":
		*post = new(VideoPost)
	case "audio":
		*post = new(AudioPost)
	case "photo":
		*post = new(PhotoPost)
	case "chat":
		*post = new(ChatPost)
	default:
		return errors.New("invalid post type")
	}

	return json.Unmarshal(b, *post)
}

func ToBasePost(post interface{}) *BasePost {
	switch post := post.(type) {
	case *TextPost:
		return &post.BasePost
	case *QuotePost:
		return &post.BasePost
	case *LinkPost:
		return &post.BasePost
	case *AnswerPost:
		return &post.BasePost
	case *VideoPost:
		return &post.BasePost
	case *AudioPost:
		return &post.BasePost
	case *PhotoPost:
		return &post.BasePost
	case *ChatPost:
		return &post.BasePost
	default:
		panic("invalid post type")
	}
}
