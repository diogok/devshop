
function vue_cart() {
  return Vue.extend({
      template:"#cart"
      ,methods: {
        rm: function(login) {
          loading();
          var that=this;
          superagent
            .del("/cart?login="+encodeURIComponent(login))
            .end(function(err,res){
                that.cart=JSON.parse(res.text);
                that.total=0;
                for(var i=0;i<that.cart.length;i++) {
                  that.total += that.cart[i].cost;
                }
                unloading();
            });
        }
      }
      ,data: function() {
        var that=this;
        setTimeout(function(){
            loading();
            superagent
              .get("/cart")
              .end(function(err,res){
                  that.cart=JSON.parse(res.text);
                  that.total=0;
                  for(var i=0;i<that.cart.length;i++) {
                    that.total += that.cart[i].cost;
                  }
                  unloading();
              });
        },500);
        return {cart:[],total:0};
      }
    });
}

