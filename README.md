# herarchy url path

gorilla/muxで階層を持つURLパスを作るにはどうしたらいいか試したやつ。

結果としては、簡単なシステムであればURLのパスはフルで書くのが良さそう。

どうしても階層化したい場合は、`PathPrefix`を使うこと。
