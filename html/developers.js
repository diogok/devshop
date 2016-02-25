
function vue_developers() {
  return Vue.extend({
      template:"#developers"
      ,methods: {
        search: function() {
          loading();
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
          this.hasPrev=true;
          this.page++;
          this.search();
        }
        ,prev: function() {
          if(this.page > 1) {
            this.hasNext=true;
            this.page--;
            this.search();
          }
        }
        ,add: function(login,cost,id) {
          loading();
          console.log(login,cost,id);
          superagent
            .post('/cart')
            .type('form')
            .send({login:login,cost:cost,id:id})
            .end(function(err,res){
                unloading();
            });
        }
      }
      ,data: function() {
        return {
          page:1,hasNext:false,hasPrev:false,hasDevs:false,query:"",developers: [ ]
        }
      }
    });
}
