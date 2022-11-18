package client

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/chromedp/chromedp"
	cdp "github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

func Setjieping(path, urlstr string, x, y int64) {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run
	var b1, b2 []byte
	if err := chromedp.Run(ctx,
		// emulate iPhone 7 landscape
		chromedp.Emulate(device.IPhone8Plus),
		chromedp.Navigate(urlstr),
		chromedp.CaptureScreenshot(&b1),

		// reset
		chromedp.Emulate(device.Reset),

		// set really large viewport
		chromedp.EmulateViewport(x, y),
		chromedp.Navigate(urlstr),
		chromedp.CaptureScreenshot(&b2),
	); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("./jt/h5_"+path, b1, 0777); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("./jt/pc_"+path, b2, 0777); err != nil {
		log.Fatal(err)
	}
}

// 截图方法
func elementScreenshot(urlstr, sel string, res *[]byte) cdp.Tasks {
	return cdp.Tasks{
		// 打开url指向的页面
		cdp.Navigate(urlstr),

		// 等待待截图的元素渲染完成
		cdp.WaitVisible(sel, cdp.ByID),
		// 也可以等待一定的时间
		//cdp.Sleep(time.Duration(3) * time.Second),

		// 执行截图
		cdp.Screenshot(sel, res, cdp.NodeVisible, cdp.ByID),
	}
}
