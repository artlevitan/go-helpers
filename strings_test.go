// Copyright 2023-2025, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.18

package helpers

import (
	"testing"
)

func TestCutString(t *testing.T) {
	type args struct {
		s      string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"", 10}, ""},
		{"2", args{"Lorem ipsum", 0}, ""},
		{"3", args{"Lorem ipsum dolor sit amet.", 10}, "Lorem ipsu"},
		{"4", args{"Lorem ipsum", 11}, "Lorem ipsum"},
		{"5", args{"Lorem ipsum", 20}, "Lorem ipsum"},
		{"6", args{"Привет, мир!", 7}, "Привет,"},
		{"7", args{"😊😃😄😁", 2}, "😊😃"},
		{"8", args{"😊😃😄😁", 3}, "😊😃😄"},
		{"9", args{"こんにちは世界", 5}, "こんにちは"},
		{"10", args{"Test long string", 4}, "Test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CutString(tt.args.s, tt.args.length); got != tt.want {
				t.Errorf("CutString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClearString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"Hello     World"}, "Hello World"},
		{"2", args{"   Hello           "}, "Hello"},
		{"3", args{"  Hello         World  "}, "Hello World"},
		{"4", args{"  Привет!     Как дела?     "}, "Привет! Как дела?"},
		{"5", args{"  Какая   чистая    строка   "}, "Какая чистая строка"},
		{"6", args{"  "}, ""},
		{"7", args{"\tHello\tWorld\t"}, "Hello World"},
		{"8", args{"\nHello \n World\n"}, "Hello World"},
		{"9", args{"Hello\n\n\nWorld"}, "Hello World"},
		{"10", args{"Привет,   мир!"}, "Привет, мир!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClearString(tt.args.s); got != tt.want {
				t.Errorf("ClearString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClearTextarea(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"Line 1\n\nLine 2\n\nLine 3"}, "Line 1\nLine 2\nLine 3"},
		{"2", args{"Line 1\r\nLine 2\r\n\r\nLine 3"}, "Line 1\nLine 2\nLine 3"},
		{"3", args{"Line 1\n    \nLine 2"}, "Line 1\nLine 2"},
		{"4", args{"\n\n\nLine 1\n\n\nLine 2"}, "Line 1\nLine 2"},
		{"5", args{"Line 1\n\n\n\nLine 2"}, "Line 1\nLine 2"},
		{"6", args{"Line 1\n \n \nLine 2"}, "Line 1\nLine 2"},
		{"7", args{"\n\nLine 1\n\nLine 2\n\n"}, "Line 1\nLine 2"},
		{"8", args{"\r\nLine 1\r\nLine 2\r\n"}, "Line 1\nLine 2"},
		{"9", args{"Line 1\nLine 2\nLine 3"}, "Line 1\nLine 2\nLine 3"},
		{"10", args{"\n \n\n \nLine 1\n \nLine 2\n \n"}, "Line 1\nLine 2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClearTextarea(tt.args.s); got != tt.want {
				t.Errorf("ClearTextarea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterLetters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"HelloWorld"}, "HelloWorld"},
		{"2", args{"H3lloW0rld"}, "HlloWrld"},
		{"3", args{"H@e#l%l^o*W(o)r_l+d!"}, "HelloWorld"},
		{"4", args{"1234567890"}, ""},
		{"5", args{"123@#$456&*7890"}, ""},
		{"6", args{"H3l@lo2 W0r^ld!"}, "HlloWrld"},
		{"7", args{"  Hello   World  "}, "HelloWorld"},
		{"8", args{"Привет Hello"}, "ПриветHello"},
		{"9", args{""}, ""},
		{"10", args{"Ёлки зелёные"}, "Ёлкизелёные"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterLetters(tt.args.s); got != tt.want {
				t.Errorf("FilterLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"1234567890"}, "1234567890"},
		{"2", args{"H3lloW0rld123"}, "30123"},
		{"3", args{"12@34#56$78"}, "12345678"},
		{"4", args{"HelloWorld"}, ""},
		{"5", args{"H3l@l#oW*r&ld!"}, "3"},
		{"6", args{"1H2e3l4l5o6W7o8r9l0d!"}, "1234567890"},
		{"7", args{" 123 456 7890 "}, "1234567890"},
		{"8", args{""}, ""},
		{"9", args{"Привет123"}, "123"},
		{"10", args{"Привет Мир"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterDigits(tt.args.s); got != tt.want {
				t.Errorf("FilterDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringLength(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{""}, 0},
		{"2", args{"hello"}, 5},
		{"3", args{"привет"}, 6},
		{"4", args{"😀"}, 1},
		{"5", args{"abc😀"}, 4},
		{"6", args{"abcdefghijklmnopqrstuvwxyz"}, 26},
		{"7", args{"こんにちは"}, 5},
		{"8", args{"你好"}, 2},
		{"9", args{"@#$%^&*()"}, 9},
		{"10", args{"𐍈😀привет"}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringLength(tt.args.s); got != tt.want {
				t.Errorf("StringLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsStringLengthInRange(t *testing.T) {
	type args struct {
		str       string
		minLength int
		maxLength int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"abc", 1, 5}, true},
		{"2", args{"abc", 4, 5}, false},
		{"3", args{"abcdef", 3, 6}, true},
		{"4", args{"abcdef", 7, 10}, false},
		{"5", args{"привет", 3, 6}, true},
		{"6", args{"привет", 7, 10}, false},
		{"7", args{"", 0, 0}, true},
		{"8", args{"a", 0, 1}, true},
		{"9", args{"a", 2, 5}, false},
		{"10", args{"𐍈", 1, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStringLengthInRange(tt.args.str, tt.args.minLength, tt.args.maxLength); got != tt.want {
				t.Errorf("IsStringLengthInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSanitizeHTML(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"<div><p>This is a <strong>test</strong> paragraph.</p></div>"}, "This is a test paragraph."},
		{"2", args{"<style>body {color: red;}</style><p>Hello, World!</p>"}, "Hello, World!"},
		{"3", args{"<script>alert('XSS');</script><p>Safe content</p>"}, "Safe content"},
		{"4", args{"<html><head><title>Title</title></head><body><h1>Hello</h1></body></html>"}, "TitleHello"},
		{"5", args{"<a href='#'>Link</a> with s"}, "Link with s"},
		{"6", args{"<div class='container'>Content inside div</div>"}, "Content inside div"},
		{"7", args{"Hello, <b>World</b>!"}, "Hello, World!"},
		{"8", args{"<blockquote>Hello, <b>World</b>!"}, "Hello, World!"},
		{"9", args{"<quietly>email me - addy in profile</quiet>"}, "email me - addy in profile"},
		{"10", args{"<p>Hello, <b onclick=alert(1337)>World</b>!</p>"}, "Hello, World!"},
		{"11", args{"<p onclick=alert(1337)>Hello, <b>World</b>!</p>"}, "Hello, World!"},
		{"12", args{`<a href="javascript:alert(1337)">foo</a>`}, "foo"},
		{"13", args{`<img src="http://example.org/foo.gif">`}, ""},
		{"14", args{`<a href="foo.html">Link s</a>`}, "Link s"},
		{"15", args{`<a href="javascript:alert(1337).html" onclick="alert(1337)">Link s</a>`}, "Link s"},
		{"16", args{`xss<a href="http://www.google.de" style="color:red;" onmouseover=alert(1) onmousemove="alert(2)" onclick=alert(3)>g<img src="http://example.org"/>oogle</a>`}, "xssgoogle"},
		{"18", args{`<img src="giraffe.gif?height=500&amp;width=500&amp;flag" />`}, ""},
		{"19", args{`<span class="foo">Hello World</span>`}, "Hello World"},
		{"20", args{`<iframe width="560" height="315" src="https://www.youtube.com/embed/lJIrF4YjHfQ" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>`}, ""},
		{"21", args{`<SCRIPT/XSS SRC="http://ha.ckers.org/xss.js"></SCRIPT>`}, ""},
		{"22", args{"<BODY onload!#$%&()*~+-_.,:;?@[/|\\]^`=alert(\"XSS\")>"}, ""},
		{"23", args{`<BODY ONLOAD=alert('XSS')>`}, ""},
		{"24", args{`test<script>alert(document.cookie)</script>`}, "test"},
		{"25", args{`<<<><<script src=http://fake-evil.ru/test.js>`}, ""},
		{"26", args{`<INPUT TYPE="IMAGE" SRC="javascript:alert('XSS');"">`}, ""},
		{"27", args{`<a onblur="alert(secret)" href="http://www.google.com">Google</a>`}, "Google"},
		{"28", args{`<BGSOUND SRC="javascript:alert('XSS');">`}, ""},
		{"29", args{`<LINK REL="stylesheet" HREF="http://ha.ckers.org/xss.css">`}, ""},
		{"30", args{`<STYLE>@import'http://ha.ckers.org/xss.css';</STYLE>`}, ""},
		{"31", args{`<META HTTP-EQUIV="refresh" CONTENT="0; URL=http://;URL=javascript:alert('XSS');">`}, ""},
		{"32", args{`<FRAMESET><FRAME SRC="javascript:alert('XSS');"></FRAMESET>`}, ""},
		{"33", args{`<TABLE BACKGROUND="javascript:alert('XSS')">`}, ""},
		{"34", args{`<DIV STYLE="background-image: url(javascript:alert('XSS'))">`}, ""},
		{"35", args{`<BASE HREF="javascript:alert('XSS');//">`}, ""},
		{"36", args{`<BaSe hReF="http://arbitrary.com/">`}, ""},
		{"37", args{`<OBJECT classid=clsid:ae24fdae-03c6-11d1-8b76-0080c744f389><param name=url value=javascript:alert('XSS')></OBJECT>`}, ""},
		{"38", args{`<EMBED SRC="data:image/svg+xml;base64,PHN2ZyB4bWxuczpzdmc9Imh0dH A6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcv MjAwMC9zdmciIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hs aW5rIiB2ZXJzaW9uPSIxLjAiIHg9IjAiIHk9IjAiIHdpZHRoPSIxOTQiIGhlaWdodD0iMjAw IiBpZD0ieHNzIj48c2NyaXB0IHR5cGU9InRleHQvZWNtYXNjcmlwdCI+YWxlcnQoIlh TUyIpOzwvc2NyaXB0Pjwvc3ZnPg==" type="image/svg+xml" AllowScriptAccess="always"></EMBED>`}, ""},
		{"39", args{`s <!-- comment -->`}, "s"},
		{"40", args{`<div>s <!--[if IE]> <!--[if gte 6]> comment <[endif]--><[endif]--></div>`}, "s comment"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SanitizeHTML(tt.args.s); got != tt.want {
				t.Errorf("SanitizeHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSanitizeHTMLWithTextarea(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"<div><p>This is a <strong>test</strong> paragraph.</p></div>"}, "This is a test paragraph."},
		{"2", args{"<style>body {color: red;}</style><p>Hello, World!</p>"}, "Hello, World!"},
		{"3", args{"<script>alert('XSS');</script><p>Safe content</p>"}, "Safe content"},
		{"4", args{"<html><head><title>Title</title></head><body><h1>Hello</h1></body></html>"}, "TitleHello"},
		{"5", args{"<a href='#'>Link</a> with s"}, "Link with s"},
		{"6", args{"<div class='container'>Content inside div</div>"}, "Content inside div"},
		{"7", args{"Hello, <b>World</b>!"}, "Hello, World!"},
		{"8", args{"<blockquote>Hello, <b>World</b>!"}, "Hello, World!"},
		{"9", args{"<quietly>email me - addy in profile</quiet>"}, "email me - addy in profile"},
		{"10", args{"<p>Hello, <b onclick=alert(1337)>World</b>!</p>"}, "Hello, World!"},
		{"11", args{"<p onclick=alert(1337)>Hello, <b>World</b>!</p>"}, "Hello, World!"},
		{"12", args{`<a href="javascript:alert(1337)">foo</a>`}, "foo"},
		{"13", args{`<img src="http://example.org/foo.gif">`}, ""},
		{"14", args{`<a href="foo.html">Link s</a>`}, "Link s"},
		{"15", args{`<a href="javascript:alert(1337).html" onclick="alert(1337)">Link s</a>`}, "Link s"},
		{"16", args{`xss<a href="http://www.google.de" style="color:red;" onmouseover=alert(1) onmousemove="alert(2)" onclick=alert(3)>g<img src="http://example.org"/>oogle</a>`}, "xssgoogle"},
		{"18", args{`<img src="giraffe.gif?height=500&amp;width=500&amp;flag" />`}, ""},
		{"19", args{`<span class="foo">Hello World</span>`}, "Hello World"},
		{"20", args{`<iframe width="560" height="315" src="https://www.youtube.com/embed/lJIrF4YjHfQ" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>`}, ""},
		{"21", args{`<SCRIPT/XSS SRC="http://ha.ckers.org/xss.js"></SCRIPT>`}, ""},
		{"22", args{"<BODY onload!#$%&()*~+-_.,:;?@[/|\\]^`=alert(\"XSS\")>"}, ""},
		{"23", args{`<BODY ONLOAD=alert('XSS')>`}, ""},
		{"24", args{`test<script>alert(document.cookie)</script>`}, "test"},
		{"25", args{`<<<><<script src=http://fake-evil.ru/test.js>`}, ""},
		{"26", args{`<INPUT TYPE="IMAGE" SRC="javascript:alert('XSS');"">`}, ""},
		{"27", args{`<a onblur="alert(secret)" href="http://www.google.com">Google</a>`}, "Google"},
		{"28", args{`<BGSOUND SRC="javascript:alert('XSS');">`}, ""},
		{"29", args{`<LINK REL="stylesheet" HREF="http://ha.ckers.org/xss.css">`}, ""},
		{"30", args{`<STYLE>@import'http://ha.ckers.org/xss.css';</STYLE>`}, ""},
		{"31", args{`<META HTTP-EQUIV="refresh" CONTENT="0; URL=http://;URL=javascript:alert('XSS');">`}, ""},
		{"32", args{`<FRAMESET><FRAME SRC="javascript:alert('XSS');"></FRAMESET>`}, ""},
		{"33", args{`<TABLE BACKGROUND="javascript:alert('XSS')">`}, ""},
		{"34", args{`<DIV STYLE="background-image: url(javascript:alert('XSS'))">`}, ""},
		{"35", args{`<BASE HREF="javascript:alert('XSS');//">`}, ""},
		{"36", args{`<BaSe hReF="http://arbitrary.com/">`}, ""},
		{"37", args{`<OBJECT classid=clsid:ae24fdae-03c6-11d1-8b76-0080c744f389><param name=url value=javascript:alert('XSS')></OBJECT>`}, ""},
		{"38", args{`<EMBED SRC="data:image/svg+xml;base64,PHN2ZyB4bWxuczpzdmc9Imh0dH A6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcv MjAwMC9zdmciIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hs aW5rIiB2ZXJzaW9uPSIxLjAiIHg9IjAiIHk9IjAiIHdpZHRoPSIxOTQiIGhlaWdodD0iMjAw IiBpZD0ieHNzIj48c2NyaXB0IHR5cGU9InRleHQvZWNtYXNjcmlwdCI+YWxlcnQoIlh TUyIpOzwvc2NyaXB0Pjwvc3ZnPg==" type="image/svg+xml" AllowScriptAccess="always"></EMBED>`}, ""},
		{"39", args{`s <!-- comment -->`}, "s"},
		{"40", args{`<div>s <!--[if IE]> <!--[if gte 6]> comment <[endif]--><[endif]--></div>`}, "s   comment"},
		{"41", args{"Line 1\n\nLine 2\n\nLine 3"}, "Line 1\nLine 2\nLine 3"},
		{"42", args{"Line 1\r\nLine 2\r\n\r\nLine 3"}, "Line 1\nLine 2\nLine 3"},
		{"43", args{"Line 1\n    \nLine 2"}, "Line 1\nLine 2"},
		{"44", args{"\n\n\nLine 1\n\n\nLine 2"}, "Line 1\nLine 2"},
		{"45", args{"Line 1\n\n\n\nLine 2"}, "Line 1\nLine 2"},
		{"46", args{"Line 1\n \n \nLine 2"}, "Line 1\nLine 2"},
		{"47", args{"\n\nLine 1\n\nLine 2\n\n"}, "Line 1\nLine 2"},
		{"48", args{"\r\nLine 1\r\nLine 2\r\n"}, "Line 1\nLine 2"},
		{"49", args{"Line 1\nLine 2\nLine 3"}, "Line 1\nLine 2\nLine 3"},
		{"50", args{"\n \n\n \nLine 1\n \nLine 2\n \n"}, "Line 1\nLine 2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SanitizeHTMLWithTextarea(tt.args.s); got != tt.want {
				t.Errorf("SanitizeHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveEmojis(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"Hello, World!"}, "Hello, World!"},
		{"2", args{"Hello, 😃 World!"}, "Hello, World!"},
		{"3", args{"Emojis 😃🚀🌍 in the middle"}, "Emojis in the middle"},
		{"4", args{"No emojis here"}, "No emojis here"},
		{"5", args{"😃 Only emojis 😃"}, "Only emojis"},
		{"6", args{"Multiple emojis 😃😃😃 in a row"}, "Multiple emojis in a row"},
		{"7", args{"Emojis at the end 😃"}, "Emojis at the end"},
		{"8", args{"😃 Emojis at the start"}, "Emojis at the start"},
		{"9", args{""}, ""},
		{"10", args{"😃"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveEmojis(tt.args.s); got != tt.want {
				t.Errorf("RemoveEmojis() = %v, want %v", got, tt.want)
			}
		})
	}
}
