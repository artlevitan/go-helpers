// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.9

package helpers

import (
	"testing"
)

func TestMD5(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{""}, "d41d8cd98f00b204e9800998ecf8427e"},
		{"2", args{"hello"}, "5d41402abc4b2a76b9719d911017c592"},
		{"3", args{"12345"}, "827ccb0eea8a706c4c34a16891f84e7b"},
		{"4", args{"The quick brown fox jumps over the lazy dog"}, "9e107d9d372bb6826bd81d3542a419d6"},
		{"5", args{"!@#$%^&*()"}, "05b28d17a7b6e7024b6e5d8cc43a8bf7"},
		{"6", args{"Привет мир"}, "79d636ccef972a9d10db69750cd53e8b"},
		{"7", args{"hello\nworld"}, "9195d0beb2a889e1be05ed6bb1954837"},
		{"8", args{"aaaaa"}, "594f803b380a41396ed63dca39503542"},
		{"9", args{"😀😃😄😁😆"}, "1100e1b35ee86b7fe1aa86840bd2cec1"},
		{"10", args{"12345678901234567890"}, "fd85e62d9beb45428771ec688418b271"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5(tt.args.s); got != tt.want {
				t.Errorf("MD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{""}, "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
		{"2", args{"hello"}, "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"},
		{"3", args{"12345"}, "8cb2237d0679ca88db6464eac60da96345513964"},
		{"4", args{"The quick brown fox jumps over the lazy dog"}, "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12"},
		{"5", args{"!@#$%^&*()"}, "bf24d65c9bb05b9b814a966940bcfa50767c8a8d"},
		{"6", args{"Привет мир"}, "6b49ee5b0520d4475013a5d6aaa085b711b3768a"},
		{"7", args{"hello\nworld"}, "7db827c10afc1719863502cf95397731b23b8bae"},
		{"8", args{"aaaaa"}, "df51e37c269aa94d38f93e537bf6e2020b21406c"},
		{"9", args{"😀😃😄😁😆"}, "68ec8840dedabfd12c0b5ea9072b70f46c7e2686"},
		{"10", args{"12345678901234567890"}, "7e0a1242bd8ef9044f27dca45f5f72ad5a1125bf"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA1(tt.args.s); got != tt.want {
				t.Errorf("SHA1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA256(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{""}, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"2", args{"hello"}, "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
		{"3", args{"12345"}, "5994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5"},
		{"4", args{"The quick brown fox jumps over the lazy dog"}, "d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592"},
		{"5", args{"!@#$%^&*()"}, "95ce789c5c9d18490972709838ca3a9719094bca3ac16332cfec0652b0236141"},
		{"6", args{"Привет мир"}, "830d1964dc8673182a40f9adebf598960d37fbe200405b249774ecfa5b465748"},
		{"7", args{"hello\nworld"}, "26c60a61d01db5836ca70fefd44a6a016620413c8ef5f259a6c5612d4f79d3b8"},
		{"8", args{"aaaaa"}, "ed968e840d10d2d313a870bc131a4e2c311d7ad09bdf32b3418147221f51a6e2"},
		{"9", args{"😀😃😄😁😆"}, "cf61553b282c801b8a7ac6db772f3d05c4d8e98faf84f350b3b7734ebaab7989"},
		{"10", args{"12345678901234567890"}, "6ed645ef0e1abea1bf1e4e935ff04f9e18d39812387f63cda3415b46240f0405"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA256(tt.args.s); got != tt.want {
				t.Errorf("SHA256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA512(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{""}, "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"},
		{"2", args{"hello"}, "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043"},
		{"3", args{"12345"}, "3627909a29c31381a071ec27f7c9ca97726182aed29a7ddd2e54353322cfb30abb9e3a6df2ac2c20fe23436311d678564d0c8d305930575f60e2d3d048184d79"},
		{"4", args{"The quick brown fox jumps over the lazy dog"}, "07e547d9586f6a73f73fbac0435ed76951218fb7d0c8d788a309d785436bbb642e93a252a954f23912547d1e8a3b5ed6e1bfd7097821233fa0538f3db854fee6"},
		{"5", args{"!@#$%^&*()"}, "138fad927473f694c3a02cca61008e52572bd19ce442f20e139b6f09157b97157fd71946fedfec2381b7e33618afe5f7c24a873ed1efe416978acfc434503614"},
		{"6", args{"Привет мир"}, "677b137ee0f83e18fb09a65001aee2fa150f28843c020eec1d4980d63c8a1d779fae69fcf0fc8030bdfe128ff04b091b220a709fbb63d51f77588839acd9f27b"},
		{"7", args{"hello\nworld"}, "c95a099794a5ef71b75704a263bec3c1f6d5d5c21f8942b82e45897321c2afb5eaa564549503869d9246ee9c912f899f052a3911733a00432dd71a77e7bae7a0"},
		{"8", args{"aaaaa"}, "f368a29b71bd201a7ef78b5df88b1361fbe83f959756d33793837a5d7b2eaf660f2f6c7e2fbace01965683c4cfafded3ff28aab34e329aa79bc81e7703f68b86"},
		{"9", args{"😀😃😄😁😆"}, "e43d260fe7eb44b2c7a42d49cbfdc8bdedcd6a84e5d9875e347ccfc653f7190b9ae5d879909017dc742530a10a8ae9c2b44a26bf9543ef10ec27bb4ccf440548"},
		{"10", args{"12345678901234567890"}, "aa3b7bdd98ec44af1f395bbd5f7f27a5cd9569d794d032747323bf4b1521fbe7725875a68b440abdf0559de5015baf873bb9c01cae63ecea93ad547a7397416e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA512(tt.args.s); got != tt.want {
				t.Errorf("SHA512() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA3_256(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{""}, "a7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a"},
		{"2", args{"hello"}, "3338be694f50c5f338814986cdf0686453a888b84f424d792af4b9202398f392"},
		{"3", args{"12345"}, "7d4e3eec80026719639ed4dba68916eb94c7a49a053e05c8f9578fe4e5a3d7ea"},
		{"4", args{"The quick brown fox jumps over the lazy dog"}, "69070dda01975c8c120c3aada1b282394e7f032fa9cf32f4cb2259a0897dfc04"},
		{"5", args{"!@#$%^&*()"}, "3168a455226ecc49217333a3632d13a568eae563fe057c6668e46322110a4670"},
		{"6", args{"Привет мир"}, "1de205a4de66fd697e9199fce67a8d59fadad61f048461cedb3116a867fd95fa"},
		{"7", args{"hello\nworld"}, "db7b8fa6a92f215757f011eff5d72c37ad272bc93d70cdd675ee9377bedd34cf"},
		{"8", args{"aaaaa"}, "04a9faa0837eb97f9d2a265ada5b1aec9fb292b3c5ab40123ef5ca9c98e55407"},
		{"9", args{"😀😃😄😁😆"}, "a41af3680bd043b680d0255aa7c71b0f6bbf71e905d53b99445c4f020fbb16ed"},
		{"10", args{"12345678901234567890"}, "672568f62f8623b49679fa5ef0aba3aac48dd19a302e9736ebb97908857e0327"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA3_256(tt.args.s); got != tt.want {
				t.Errorf("SHA3_256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA3_512(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{""}, "a69f73cca23a9ac5c8b567dc185a756e97c982164fe25859e0d1dcc1475c80a615b2123af1f5f94c11e3e9402c3ac558f500199d95b6d3e301758586281dcd26"},
		{"2", args{"hello"}, "75d527c368f2efe848ecf6b073a36767800805e9eef2b1857d5f984f036eb6df891d75f72d9b154518c1cd58835286d1da9a38deba3de98b5a53e5ed78a84976"},
		{"3", args{"12345"}, "0a2a1719bf3ce682afdbedf3b23857818d526efbe7fcb372b31347c26239a0f916c398b7ad8dd0ee76e8e388604d0b0f925d5e913ad2d3165b9b35b3844cd5e6"},
		{"4", args{"The quick brown fox jumps over the lazy dog"}, "01dedd5de4ef14642445ba5f5b97c15e47b9ad931326e4b0727cd94cefc44fff23f07bf543139939b49128caf436dc1bdee54fcb24023a08d9403f9b4bf0d450"},
		{"5", args{"!@#$%^&*()"}, "fbbcb3e21184dd4061de0b85c4756f74e36a361125733e2c7470232fc66f71c902d1e6ff7eb60cfe6b47e8b72e1429b4ff21de0fa150a2b3e8d8e29e264d56ab"},
		{"6", args{"Привет мир"}, "6f826c212e5ef27ab815fb8f2fd90675aefffee363fc5ce4fa0081af62a8fe7421e3177e2bf3b1ac0eee14e8ec6ebff7b9e46ac6df0cbc4601c9fd71e8fa368f"},
		{"7", args{"hello\nworld"}, "9418d957d7a7486841add1139f17e7b12a1f0503ecb6b834d6a14f2a74e39ac75277ba5358df094dbce4474952ebaf1e64a599476c08b54b3d53b2702c37338c"},
		{"8", args{"aaaaa"}, "384337abeaa3a24884ba6ce6df7e7c533569091f89f102a940ac19242e4947ac41d80c5e2fb4babc825113d2c06c5e44a39c9da3ca4d3fb8cf5969c2def21c7f"},
		{"9", args{"😀😃😄😁😆"}, "7aed3603afba0fa526818b19da678ccf38c2d8da93c299c1081f06426bbd32fe61ddd540a01c4cb4d867290c1029114f681b142b63abda8a0dc1ed5a681e8230"},
		{"10", args{"12345678901234567890"}, "8e61e13f551c5aa74282aecafdae5d89e8508520c68c85e5271139fe08de96be4709825cdfafa736753c052df509c53cd9e6e6e452c73c85538cba9fd35bc598"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA3_512(tt.args.s); got != tt.want {
				t.Errorf("SHA3_512() = %v, want %v", got, tt.want)
			}
		})
	}
}
