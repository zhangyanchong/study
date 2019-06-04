"use strict";
var page = require('webpage').create(),
    system = require('system'),
    address, output, size;
address = system.args[1];
output = system.args[2];

page.open(address, function (status) {
    if (status !== 'success') {
        console.log('Unable to load the address!');
        phantom.exit(1);
    } else {

        window.setTimeout(function () {
             page.evaluate(function() {
                 $(".zycdel").remove();  // 删除所有不用的js
             });
            console.log(page.content);
            //  var p = page.evaluate(function () {
            //
            //      return document.getElementsByTagName('html')[0].innerHTML
            //  });
            phantom.exit();
        }, 10000);
    }
});