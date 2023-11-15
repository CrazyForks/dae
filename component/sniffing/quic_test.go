package sniffing

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/daeuniverse/dae/component/sniffing/internal/quicutils"
	"github.com/sirupsen/logrus"
)

var QuicStream2_2, _ = hex.DecodeString("cc0000000108e8da6ed9f385c987000044d026f109c2764c22f0ea2656550ea03e832d0ed5113eff115f2a057f77655cf5bbbb69fc98f7f70a3f407e0d94f37960c5ba5bd95a2df75f6f25020c2f2f21ddf9db5266bb4293991d58efec945468a820c61b743ca4b73663c3adcda58dee75607c5465e255b58477069a928687789c18c2ccb53911a47d64b83d5b58398ee4fd58f4f88f78788d5594218730cab9db3bac2fbfb947f2cb4eafb5e2964fce361042c622dfa7130afaf0e9d391ffc3aba2f5ee2f5c4d0dfaae0d71db2b3d7fab6dbccbb63d7961ddab55711d5a1beacf00ce5a82030a2c79c4ea65a2762f3b8e5f8fec8f6963b1a42c0f8a8d863225b2d6e7a15e9758e43095459e3d7ff88dc276605452b10de95a8795fe9952eb0b1eb200465ca9b00f98e2c4ad6a2a2e2bff2e2430438241525e1d16d5423c2262134a97056b7e86d5eb7eb2ac546086a3b8d7a97bc2263fa9a8b46f4b7d31cad63762c17a653b89593434aecf7a5e8fc169cfb5aa4a47e78ee817e115feceb9b68b29da6e15c647b7528980fb7cdc7c9ca660871228d0367f030f658d19ddddefe55908a2ec4ef5f5d89ec5aebee33f88a116c2857f7d1a2fd98321f28468a93938da406a68e4e660f0668fe49118812d5264073f28a8aa800c5970ef3f6fb4f0e9e4e48510700a5465c92886c50f2c6af570075f29f6a80636171f73d91864583d2d199e39b18623ee0cb489b449838bd9f7cd67ccc3e38f1b5a3ce08814f979f94db45cdcfa39a475e3efc4847def8e8e4c707a88d2f486fc85e10910ab0f1bbeb40468af777ff2bb0e655f1a006cde0d2e2ae036dafe60f110e859543699e0c9aa47eefa53d792b3cbcfa11ea1d3b55d3629de0345517d47f4e4c801104b81710ad28cd8611e150a1fc32160cb784cfcfdd908052cd43969b27929013edd2b0f3cd914590a32b2f99d4fc88873838b6fa0ec1450adb95f395988998801e85319fa448925ba767e3191df2b5b0983990beb4127216c93291a94463b453a4972c9a974742b0b22c935f4235c350120b6cf8296fc6d3c2812f74a17acf334e3c34ff9988f980e0cfff737a8b1a03508f47d8bf3748fbb5bd5ad7f1f47120c3a33822612f3a614aae7fe536b73db814aa4aac4b685aa1e7357309cf921b931113624881ce764feeff3292d2d794c6fa76529f3da8e6327e8f28aafe8b675a80ae3f478c65f1bf8fd7f2b140fea130dfa55982f0b0fcd61b42c8b2ea27a2b8bb44511eb44c1416ac16698f0ddb739e3d773f2afdd35bcfed0ffd7966aa3e727f8f08d02cab8d034a7ae363e42c9089901ddee147c98a856df4e5dcfeeb2f72e9edb12da513f32d99e1c653f4503e9a7f7fee1f4724ce9d6d530485362d993cb3bc4faff683327a02aee6f004bd9f98a8a4841091d48f5cd27af46431c66e68007750be57361e293650a0ae9fc9fa82ddf4483663c9805dc6e4a9b43529c0b2267cc3c0fb9084378acbda4962150a73e0c1b5aef6e40538d2630d8dbc2b084f9a53079cc73484906b7ad4a5021f280baf276a01b0fcea57d5c4284364f4d795645fc7bd8bb7d00021af924b75829e8a936e153676a182803537a23c76fee7c881e8063751ca0f5a585481b9077e9593734f9997e78b79ba38f6e13a1b631106a2ceddafdf51110b8bf07ec9337024355088d0bb3de2d46a03d3e3e7362b8b815613e36d746e5a9992f8e62ad5257e5798bd49b1a62717f02151b75a18e051df1292191d4")
var QuicStream2_1, _ = hex.DecodeString("ce0000000108e8da6ed9f385c987000044d0f34f94dcc26b99261ea264742abe4e552a146e16e89e4b7ef0ab3d6f3a34227b59742e4ba83a1e18cea494d2f67e469be4a7ff01334b151e9b7ca63b53735008eecc1f5c618419982292eca5731bb163ba81c1300e0bb99f2536d89ab0faf2dbd37ebfdb3d71f7343296a2190914bda556b8f9ccf5219964eb3cd373966fcfaca8a4735fb59fbaf69bbbdfc3a81b11570bb81fd3f5ef780fb7036e0666b997b0f4ed3305b68eafa1a99b3c8a6a2142ad9fe1e6b0a0eade6ace92b57416d4bf68fa2e9295bfc22757b0542ce91c8af3f547ef0ad385788db230a50158a0009fd95a7e8ee6e0dd11d6f9a906cbe8117e85bd507cdbd8f1a5a6cabf2617de7227d1ae8a8c6086b8ec325df90c0e16b37b4ed0ce617a00c7598a21924a19aec1b08c31b69430b23eefbe555ca2433431d28a4ffec548e463e8e6363b6b4fe9b8477c686c393571273c30b2e1785261faa0fd6f560c12418b27cd0491e013db5a8b3294e01a46a6e4c6b52e32756ab4be6f4ebc886c0c472d63f117ce30115182a97f1308c7f28989ce301cabced825154b0f4fa3bf4a55ce2f384ff11d9cbc0460d69db363664f92dc014bdb771b9b1e1ab6672c6da71c90aa514dcdc3a4ce45298bf9e5a395ebac3dff2a738c4b4690ee06fdab572a277addac7035d94afe794df05da75a56c79c37f42de1d727dc65e3060d9331e2fc82de2d7cef6cb9ae46f648b9930593975c35960b24deb770d5ee4332f8f57a05503399ca7bfdf7207f66a0f73d6b53269a944d5a3043b225adddfdd29d20ea8f500bb09ea3bb724083dd29ea8839e8192c4360ba3c5a6db0d695af5d357d6c4ed94aa28305033629201689764189774bbd4f0ae41b878b8f29a0fe0e124075ea08c5054871506a05be2f90e9ec0c2db48c0780580312e9ff4071054386e4206841f575f7ca06c228f7ee11e2333d08652b9b4f0b97f473a46a3d79c4f9a3416fb20fdbd88cacfa36f06fe1d73618195c6f0bf759a77c6a16b7e271c6cdb672ea53f6edfac860fcaf03313564abde1f66bca441d844d289a9e1025711c284f2c7c805353f2a89e9aeb52e3f452e879f0fafcdc0b48a0676afcf617a85037d991762664f6db64847eff2308447c4e8ea6688838bb7237a5fdfe0f1695afaa0bbb821b0004585adf151b029bd3458e28ba49dfc17eef1d2dd14ccda88d0848d4cd36d33cc5bab173c2448785ec1bdabc8873c904b95d7847d1b89857f2c7e078c6e2eb96029aa91c077e0efcf7b2ed2f30c7abc12189627793c7870dc0e70342cc27402ee1d6dec5ceea0ca06159002ea14a20c63b85689ed1840f404e46cb83d91c5e02f3ed938462364d3349f689310234083f7044e4b338ac54bed94530640d684c9688651b915d8c8895ef0f05f376292871b589751ac5b233e3d85572bb0c11bbbe91cc49a4ef0422f2676a2f3cc62bc88dbb7acf03cb5e847e976bfca6a90b9cee743ea77be5472ef162ff101c6873043df94c53c252840fd6a2662018f0897a06cd215997d6050917876500796fef718957212c773c39d1c7b839931af1e7dfae6e2c1d2251e78896521bb35b20057bad77df85aaed90288c17edb081398815e47239aeb77293a02a61a5125109fc3953593233fa83c17770a815fad7831c1b8647c6089ec621ee774a12a714def498d4335d0bb8a4a6a3dddead8ddb1176f58218477d55317df88cd2ca5a06b72679cf2ff7253ebd76a5ed3")
var QuicStream3, _ = hex.DecodeString("c00000000110787cb250e5ebaa3070534ac6f568006c14376bb3d77569ef83965513f7ab60499d3d6fe8cd00411e61c97af492e1c220194c2460a093505250315e811506fda1a54b7b6bfc85e18d997db284c578a4c4576258c92176200b5f85d40b28734880c8c01a9e9d5944b17568a24e112e966bf0ee955981635f0dde48e0d176f8492708a4436a53a4794a29dd8b020521824823db71bb6a4266baaf9364a2268cf87ee1dd9a543c9268c3d7ef6726e9bdea6f38d615b9ba08b3a290a22ebc1fcd9093bde5098c3c0d6151ab1e30243d21906a88e8d248a55a2c4d282e309fced134e4d13d9d2ef49325a2741824b14f1a018cfed76d0de5b6cd2881c0c708bbcca59cff5cb60ad7b9a2909b1afb4efe0b358ba098b6b2a598da1f9d23accdab814f524c1e1e0d86d3c1e4199b358a5dad8eacfe6d5d1cf431a44129538177824ed150650d97631d4d00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")

func dumpCryptos(t *testing.T, cryptos []*quicutils.CryptoFrameOffset) {
	var b strings.Builder
	for _, c := range cryptos {
		b.WriteString(fmt.Sprintf("Offset %v; length: %v:\n", c.UpperAppOffset, len(c.Data)))
		b.WriteString(fmt.Sprintf("Dump:\n%v\n", hex.Dump(c.Data)))
	}
	t.Log(b.String())
}

func TestQuicReassemble(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	sniffer := NewPacketSniffer(QuicStream2_1, 300*time.Millisecond)
	d, err := sniffer.SniffQuic()
	if err != nil {
		if sniffer.NeedMore() {
			sniffer.AppendData(QuicStream2_2)
			d, err = sniffer.SniffQuic()
		} else {
			t.Fatal(err)
		}
	}
	dumpCryptos(t, sniffer.quicCryptos)
	if err != nil {
		t.Fatal(err)
	}
	if d == "" {
		t.Fatal("domain is empty")
	}
	t.Log(d)
}
func TestQuic(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	sniffer := NewPacketSniffer(QuicStream3, 300*time.Millisecond)
	d, err := sniffer.SniffQuic()
	if err != nil {
		dumpCryptos(t, sniffer.quicCryptos)
		if sniffer.NeedMore() {
			t.Fatal("need more")
		} else {
			t.Fatal(err)
		}
	}
	dumpCryptos(t, sniffer.quicCryptos)
	if err != nil {
		t.Fatal(err)
	}
	if d == "" {
		t.Fatal("domain is empty")
	}
	t.Log(d)
}
