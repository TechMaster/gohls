package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"github.com/TechMaster/gohls/hls"
	"net/http"
	"github.com/kataras/iris"
)

type streamCmd struct {
	port    int
	homeDir string
}

func (*streamCmd) Name() string     { return "stream" }
func (*streamCmd) Synopsis() string { return "Stream Video" }
func (*streamCmd) Usage() string {
	return `strean <path to videos>:
  stream videos in path as HTTP
`
}

func (p *streamCmd) SetFlags(f *flag.FlagSet) {
	f.IntVar(&p.port, "port", 8080, "Listening port")
}

func (p *streamCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	app := iris.New()

	// Generate variables and paths
	var port = p.port
	var videoDir = setVideoDir(f)

	// Setup HTTP server
	/* old code
	http.Handle("/", http.RedirectHandler("/ui/", 302))
	http.Handle("/ui/assets/", http.StripPrefix("/ui/assets", &assetHandler{}))
	http.Handle("/ui/", NewDebugHandlerWrapper(http.StripPrefix("/ui/", NewSingleAssetHandler("/index.html"))))
	*/

	app.Get("/", func(ctx iris.Context) {
		ctx.Redirect("/index2.html", 302)
	})

	app.StaticEmbeddedGzip("/", "./ui/build", GzipAsset, GzipAssetNames)





	http.Handle("/list/", NewDebugHandlerWrapper(http.StripPrefix("/list/", hls.NewListHandler(videoDir))))
	http.Handle("/frame/", NewDebugHandlerWrapper(http.StripPrefix("/frame/", hls.NewFrameHandler(videoDir))))
	http.Handle("/playlist/", NewDebugHandlerWrapper(http.StripPrefix("/playlist/", hls.NewPlaylistHandler(videoDir))))
	http.Handle("/segments/", NewDebugHandlerWrapper(http.StripPrefix("/segments/", hls.NewStreamHandler(videoDir))))
	http.Handle("/download/", NewDebugHandlerWrapper(http.StripPrefix("/download/", NewDownloadHandler(videoDir))))

	// Dump information to user
	fmt.Printf("Path to ffmpeg executable: %v\n", hls.FFMPEGPath)
	fmt.Printf("Path to ffprobe executable: %v\n", hls.FFProbePath)
	fmt.Printf("Home directory: %v/\n", hls.HomeDir)
	fmt.Printf("Serving videos in %v\n", videoDir)
	fmt.Printf("Visit http://localhost:%v/\n", port)

	/* old code
	if herr := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); herr != nil {
		fmt.Printf("Error listening %v", herr)
	}*/

	app.Run(iris.Addr(":8080"))
	return subcommands.ExitSuccess
}