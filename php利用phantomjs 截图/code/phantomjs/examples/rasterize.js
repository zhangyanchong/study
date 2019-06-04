"use strict";
var page = require('webpage').create(),
    system = require('system'),
    address, output, size;
    address = system.args[1];
    output = system.args[2];
	var webWidth = system.args[3];   //自己做的优化  页面的真实宽度 需要传递过来       （如果有问题请查看备份的代码）
//  console.log(system.args.length); 
   if(system.args.length > 4){
        try {
            var  webCookie=system.args[4];
            var  webDomain=system.args[5];
                  webCookie=decodeURIComponent(webCookie);
            //  console.log(webCookie);
               //  console.log(webDomain);
              var   webCookieObj=JSON.parse(webCookie);
            // phantom.exit();

           phantom.clearCookies();
             for (var b in  webCookieObj) {
               //     console.log(b);
               //  console.log(webCookieObj[b]);  phantom.exit();
                phantom.addCookie({
                    'name': b, /* required property */
                    'value': webCookieObj[b], /* required property */
                    'domain': webDomain,
                 //   'path': '/', /* required property */
             //       'httponly': true,
              //      'secure': false,
             //       'expires': (new Date()).getTime() + (1000 * 60 * 60)   /* <-- expires in 1 hour */
                });
               
             }
        }catch (e){
            console.log(e);
        }
    }



    page.viewportSize = { width: webWidth, height: 600 };
    page.open(address, function (status) {
        if (status !== 'success') {
            console.log('Unable to load the address!');
            phantom.exit(1);
        } else {
            window.setTimeout(function () {
                page.render(output);
                phantom.exit();
            }, 5000);
        }
    });

