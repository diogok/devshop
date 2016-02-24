
function main() {

  var index = Vue.extend({template:"#index"});
  var cart = Vue.extend({template:"#cart"});

  var app = Vue.extend({});
  var router = new VueRouter();

  router.map({
     '/':{component: index}
    ,'/cart':{component: cart}
  });

  router.afterEach(function(){
      componentHandler.upgradeDom();
  });

  router.start(app,".inner-content");

}

window.onload = main;
