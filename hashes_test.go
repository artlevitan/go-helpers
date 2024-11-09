// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.13

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

func TestSHA3_224(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{""}, "6b4e03423667dbb73b6e15454f0eb1abd4597f9a1b078e3f5b5a6bc7"},
		{"2", args{"hello"}, "b87f88c72702fff1748e58b87e9141a42c0dbedc29a78cb0d4a5cd81"},
		{"3", args{"12345"}, "94cc697550f5c7399d179e206cf1e7bf90e17de8a87ff0f9368ec839"},
		{"4", args{"The quick brown fox jumps over the lazy dog"}, "d15dadceaa4d5d7bb3b48f446421d542e08ad8887305e28d58335795"},
		{"5", args{"!@#$%^&*()"}, "20ffee9e8da8fdd7d1dc8bcded0bb585affc7d5064d7056242947054"},
		{"6", args{"Привет мир"}, "f64dd400850e74ee3cc6f6e1f7d389861b7a441bcb9dffff2ae7bfa1"},
		{"7", args{"hello\nworld"}, "60e6e9254e935392e1d7b610a70f52e38db09ce3008ba0b5887a27a4"},
		{"8", args{"aaaaa"}, "ac341194ec843e67c48471557058aa6a10345003fe25089f10728b62"},
		{"9", args{"😀😃😄😁😆"}, "bca215bae62041c9b147d023d95130d6c59066c4e59b416e7dcccd32"},
		{"10", args{"12345678901234567890"}, "17358d187a0d65739e3e2f9b694f6412591a2e6d4e2337e01dd3bab7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA3_224(tt.args.s); got != tt.want {
				t.Errorf("SHA3_224() = %v, want %v", got, tt.want)
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

func TestSHA3_384(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{""}, "0c63a75b845e4f7d01107d852e4c2485c51a50aaaa94fc61995e71bbee983a2ac3713831264adb47fb6bd1e058d5f004"},
		{"2", args{"hello"}, "720aea11019ef06440fbf05d87aa24680a2153df3907b23631e7177ce620fa1330ff07c0fddee54699a4c3ee0ee9d887"},
		{"3", args{"12345"}, "161609f9697539edd5e03b6f5bfd1735f5c6037e0b00027c45a80386d5ebdcd3eb4bde062710914c7f37bd45f1c8021d"},
		{"4", args{"The quick brown fox jumps over the lazy dog"}, "7063465e08a93bce31cd89d2e3ca8f602498696e253592ed26f07bf7e703cf328581e1471a7ba7ab119b1a9ebdf8be41"},
		{"5", args{"!@#$%^&*()"}, "15fdc66ff959f7eba22bfe7abcaf316037cdaee1c227892bcd3a13a2e5d891cfc861cf6b33998bc75359bf3598569549"},
		{"6", args{"Привет мир"}, "1b8208671fdc0372f98f6f8a84d0341883468fcdad68e1811c9d875ec2e8da303c529a4e24d62a931daa43cdaad958da"},
		{"7", args{"hello\nworld"}, "979a690013fc81c611c3e3e63d90903717b38db3aac746246eb721f269df8424ec8fc9d705afb3bc36f2ac9e9aa6f12c"},
		{"8", args{"aaaaa"}, "48fb8f2fe35e680dc4da73dfee3ed26f7def817b9900ed247da3ce35268a6355be96a016460bbe24e2d8fdbe058908f9"},
		{"9", args{"😀😃😄😁😆"}, "cb392afeb0eacd2e81625ff67fa10fc8e752b86c208d9ea2d87b33a712d1a43a4ff7e6b35b89f5b029d14f162dbd3d17"},
		{"10", args{"12345678901234567890"}, "566c0e5584e50132c739a7070ef1b169bee92e91ce5ece1050685f31902e33ff3f1100d506ef0e4e1bd44eca3600000b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA3_384(tt.args.s); got != tt.want {
				t.Errorf("SHA3_384() = %v, want %v", got, tt.want)
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

func TestSHAKE128(t *testing.T) {
	type args struct {
		s            string
		outputLength int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"", 32}, "7f9c2ba4e88f827d616045507605853ed73b8093f6efbc88eb1a6eacfa66ef26"},
		{"2", args{"hello", 32}, "8eb4b6a932f280335ee1a279f8c208a349e7bc65daf831d3021c213825292463"},
		{"3", args{"12345", 32}, "406a6ebf1a5aaecf405f9bb5a51ecc3777f6e4c83755a641b14049400b78999e"},
		{"4", args{"The quick brown fox jumps over the lazy dog", 32}, "f4202e3c5852f9182a0430fd8144f0a74b95e7417ecae17db0f8cfeed0e3e66e"},
		{"5", args{"!@#$%^&*()", 32}, "f87b044fe7145e2b29d581a0bf18774902a19022c0a0c19f09c303f94f61b2dc"},
		{"6", args{"Привет мир", 32}, "633b170c20b75262c5ab82b1595b8dc6410d74dbf6961937ca7aab3337699c59"},
		{"7", args{"hello\nworld", 32}, "c3d16fbe6a5997401fcf3c2a2bba265ba79e9b2373923446115334f58295033f"},
		{"8", args{"aaaaa", 32}, "5bd8c1f8e2bf8d184d085fec641a82c0d72f295d3b1261960f2fc1ec7a36004e"},
		{"9", args{"😀😃😄😁😆", 32}, "1b2f7b70981052a6b6da773b3237cf81ed42dd27ac92b1b52def8446a5e072a1"},
		{"10", args{"12345678901234567890", 32}, "0a6ffcb1112f7ba7da7fc7d2024c20e61d0cb224879db482aea4cbf644b94769"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHAKE128(tt.args.s, tt.args.outputLength); got != tt.want {
				t.Errorf("SHAKE128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHAKE256(t *testing.T) {
	type args struct {
		s            string
		outputLength int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"", 64}, "46b9dd2b0ba88d13233b3feb743eeb243fcd52ea62b81b82b50c27646ed5762fd75dc4ddd8c0f200cb05019d67b592f6fc821c49479ab48640292eacb3b7c4be"},
		{"2", args{"hello", 64}, "1234075ae4a1e77316cf2d8000974581a343b9ebbca7e3d1db83394c30f221626f594e4f0de63902349a5ea5781213215813919f92a4d86d127466e3d07e8be3"},
		{"3", args{"12345", 64}, "608e70ec1937c4b0baf8c471f679485b8a769c27249ba85371fc252f8a13de30b303e6776ad14de7a6766f64ffb4cfd7d50938e97017a0b6163cfd9421548d11"},
		{"4", args{"The quick brown fox jumps over the lazy dog", 64}, "2f671343d9b2e1604dc9dcf0753e5fe15c7c64a0d283cbbf722d411a0e36f6ca1d01d1369a23539cd80f7c054b6e5daf9c962cad5b8ed5bd11998b40d5734442"},
		{"5", args{"!@#$%^&*()", 64}, "7032164a32b6ed65a42c0add2ee8b2ca1f1070a75a41484bec80a996da7ac5867b887ae9febe036ebff0618a8b3f7a913ef2028b4fcce7995bf850c5a132c58a"},
		{"6", args{"Привет мир", 64}, "8b3f17e3b2249469e17fd6105b5cf5424afd58e255f899ee339cfb074ae00cfa7838d4475082c4a391422afaf131961285c66102db2e28792c1e88b59350d7c8"},
		{"7", args{"hello\nworld", 64}, "3b6fab63e057c229e295edf20d4d6a05783919f0bec0797b50be86b0877d9bd49cdee18fe842e6a0711ea48fa1293c556b658c3727f2cd0ebe24291ef1081b07"},
		{"8", args{"aaaaa", 64}, "c168295cbb212a077e3c3f2b6757dddf9c6bb57851fb71acc73c3f5e0dc568c7b2e61460b392b89b54aaa7ca1893c8c9f9133b16c297eb8bf9e9c934d0306f00"},
		{"9", args{"😀😃😄😁😆", 64}, "dd1427753b27303f377d5b5b8dc5e050e1868c4db9ab0f05ea53dd2e9c81905b45b9b6ceb091a3eed583f6db4265d3267c8c7529294418f23093a14874e55446"},
		{"10", args{"12345678901234567890", 64}, "e83230313cfe1c52451ddb1194a4fd77177c22852a813d3b829734486c011c48e356fcb6c345361e6857a0bd35d4828c4c5844442ee5be7abac573fd0ecb3536"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHAKE256(tt.args.s, tt.args.outputLength); got != tt.want {
				t.Errorf("SHAKE256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBLAKE2b_256(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{""}, "0e5751c026e543b2e8ab2eb06099daa1d1e5df47778f7787faab45cdf12fe3a8"},
		{"2", args{"hello"}, "324dcf027dd4a30a932c441f365a25e86b173defa4b8e58948253471b81b72cf"},
		{"3", args{"12345"}, "0c79cdda0b4fca6014a1a171b608f17d4f373bc7f1187c3cf87fcffa490fc638"},
		{"4", args{"The quick brown fox jumps over the lazy dog"}, "01718cec35cd3d796dd00020e0bfecb473ad23457d063b75eff29c0ffa2e58a9"},
		{"5", args{"!@#$%^&*()"}, "058b5307ad72c25c68edb177e6d890112c023373124994889b74cc1c51508b72"},
		{"6", args{"Привет мир"}, "2714b403c8f60e2ffea750140534a7da87e7379cfc398e9d73bfebf5a56d2716"},
		{"7", args{"hello\nworld"}, "5fadc6e0432a4c6f929c823db2c979d7cf4a06f731cf4c07988ae59e9adb88a1"},
		{"8", args{"aaaaa"}, "bd9a50fa60ab6ecedb81dee4d80e715156dd2abde11c2fa27ea69e55a8ef1779"},
		{"9", args{"😀😃😄😁😆"}, "8e2e05b07515c7020dbecd7fd70181086f4320217ea754ceb0dd88b3fe89758a"},
		{"10", args{"12345678901234567890"}, "bce9c1dea3d68eab351446cb42420289b0fb87c525ff250c5d2801b0ca3d08b8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BLAKE2b_256(tt.args.s); got != tt.want {
				t.Errorf("BLAKE2b_256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBLAKE2b_512(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{""}, "786a02f742015903c6c6fd852552d272912f4740e15847618a86e217f71f5419d25e1031afee585313896444934eb04b903a685b1448b755d56f701afe9be2ce"},
		{"2", args{"hello"}, "e4cfa39a3d37be31c59609e807970799caa68a19bfaa15135f165085e01d41a65ba1e1b146aeb6bd0092b49eac214c103ccfa3a365954bbbe52f74a2b3620c94"},
		{"3", args{"12345"}, "8b28f613fa1ccdb1d303704839a0bb196424f425badfa4e4f43808f6812b6bcc0ae43374383bb6e46294d08155a64acbad92084387c73f696f00368ea106ebb4"},
		{"4", args{"The quick brown fox jumps over the lazy dog"}, "a8add4bdddfd93e4877d2746e62817b116364a1fa7bc148d95090bc7333b3673f82401cf7aa2e4cb1ecd90296e3f14cb5413f8ed77be73045b13914cdcd6a918"},
		{"5", args{"!@#$%^&*()"}, "3eafdd8f107221b7900235627e0ab999140e93ca249c65b678737b2d66262bb0abae444dbc2550f56fe359ccb75f417334c8676f1c1fb2a0990cf5bd28046a05"},
		{"6", args{"Привет мир"}, "b3d8d6ee8cf3dc6f6d7620b884e370e407cdb0e4b761ae9b5f8db0daf9776397ff4413d44ffdb2d33f017d86fcd33774f0b4ef08a5d5674a681a2fba89de8fe4"},
		{"7", args{"hello\nworld"}, "2da1ac46f8a41088ce1f9bb3d6b66bd0f219c1bd06df48bbac2fc26537bf19deac9ce08e16ec97efc08b94e02237c4d364edcf87f1c76bccaee5edc6ed309088"},
		{"8", args{"aaaaa"}, "88f831a0766167e6258c61c750ab0bc4051a2eb433397178367beaa76e97e4077e9d75d8d908ea57e19d9f3e07bd3757bdefd0f5d888d5fa501372f5475720ed"},
		{"9", args{"😀😃😄😁😆"}, "cb65e32d6c619ccb0a57905e8b66bf166c06bb5a90fb4a0e77d5690cf7b2f225e0f0ccbe4ab07d21743b838e904874d92c72848cfc0aa5c8f1139e2cf4d5fc90"},
		{"10", args{"12345678901234567890"}, "99c28b35df7614d1dea68b00c22ee913afd8d03406312355205424e29d0a1cd612d3713de20797732c71640e253dacd21f5ed99752f165160f68c43a3ae17ce9"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BLAKE2b_512(tt.args.s); got != tt.want {
				t.Errorf("BLAKE2b_512() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBLAKE2s_256(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{""}, "69217a3079908094e11121d042354a7c1f55b6482ca1a51e1b250dfd1ed0eef9"},
		{"2", args{"hello"}, "19213bacc58dee6dbde3ceb9a47cbb330b3d86f8cca8997eb00be456f140ca25"},
		{"3", args{"12345"}, "a076a699190673026fe44f7b523d321fcae79e70945007bdb1c86295a11c4135"},
		{"4", args{"The quick brown fox jumps over the lazy dog"}, "606beeec743ccbeff6cbcdf5d5302aa855c256c29b88c8ed331ea1a6bf3c8812"},
		{"5", args{"!@#$%^&*()"}, "62e01964bafc62b1d4c439ffb369f35e53230794a62ab079029caa8bbe5c6ef0"},
		{"6", args{"Привет мир"}, "8cd4785d0243a997f2950c221943bfb273f0b5587f7e5b50300d92fa8f266014"},
		{"7", args{"hello\nworld"}, "d2f2dce333ba6323bf87fe0720662211b0a32b644023539f490c909c4ba4d549"},
		{"8", args{"aaaaa"}, "97e2f2703b00d468b3b32ce91305570c6e67c08b16c98e52fdc11160515c7c69"},
		{"9", args{"😀😃😄😁😆"}, "81b1ed36e0157eaac3742b543557fb0b370ff54370f51eca936e2f5db5459f2d"},
		{"10", args{"12345678901234567890"}, "1f974c063969eff54ecdaa9683caeb14bed0908333d86bfe268e46d35e4569f7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BLAKE2s_256(tt.args.s); got != tt.want {
				t.Errorf("BLAKE2s_256() = %v, want %v", got, tt.want)
			}
		})
	}
}
