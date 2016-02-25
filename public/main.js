
function main() {
  var index = Vue.extend({template:"#index"});

  var cart = Vue.extend({template:"#cart"});

  var developers = Vue.extend({
      template:"#developers"
      ,methods: {
        search: function() {
          loading();
          event.preventDefault();
          var that=this;
          superagent
            .get('/developers?page='+this.page+'&query='+encodeURIComponent(this.query))
            .end(function(err,res){
              that.developers = JSON.parse(res.text);
              if(that.developers.length >= 5) {
                that.hasNext=true;
              }
              if(that.developers.length != 0) {
                that.hasDevs=true;
              } else{
                that.hasDevs=false;
              }
              unloading();
            });
        }
        ,next: function(){
          event.preventDefault();
          this.hasPrev=true;
          this.page++;
          this.search();
        }
        ,prev: function() {
          event.preventDefault();
          if(this.page > 1) {
            this.hasNext=true;
            this.page--;
            this.search();
          }
        }
        ,add: function() {
          event.preventDefault();
          loading();
        }
      }
      ,data: function() {
        return {
          page:1,hasNext:false,hasPrev:false,hasDevs:false,query:"",developers: [ ]
        }
      }
    });

  var app = Vue.extend({});
  var router = new VueRouter();

  router.map({
     '/':{component: index}
    ,'/cart':{component: cart}
    ,'/developers':{component: developers}
  });

  router.afterEach(function(){
      componentHandler.upgradeDom();
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
