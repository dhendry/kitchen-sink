var Test2 = new Phaser.Class({

    Extends: Phaser.Scene,

    initialize:

        function Test2 (config)
        {
            Phaser.Scene.call(this, config)
        },

    preload: function ()
    {
        this.load.image('face', 'assets/pics/bw-face.png');
    },

    create: function ()
    {
        this.face = this.add.image(400, 300, 'face');
    }

});
