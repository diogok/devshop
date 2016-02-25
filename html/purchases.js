
function vue_purchases() {
  return Vue.extend({
      template:"#purchases"
      ,data: function() {
        var that=this;
        setTimeout(function(){
            loading();
            superagent
              .get("/purchase")
              .end(function(err,res){
                  that.purchases=JSON.parse(res.text);
                  unloading();
              });
        },500);
        return {purchases:[]};
      }
    });
}


