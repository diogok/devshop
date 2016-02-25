
function main() {
  var index = Vue.extend({template:"#index"});

  var developers = vue_developers();
  var cart = vue_cart();

  var app = Vue.extend({});
  var router = new VueRouter();

  router.map({
     '/':{component: index}
    ,'/cart':{component: cart}
    ,'/developers':{component: developers}
  });

  router.afterEach(function(){
    setTimeout(function(){
      componentHandler.upgradeDom();
    },700);
  });

  router.start(app,".inner-content");

}

function loading() {
  document.querySelector(".loading").classList.add('is-active');
}

function unloading() {
  document.querySelector(".loading").classList.remove('is-active');
}

window.onload = main;
