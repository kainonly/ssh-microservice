<?php
declare(strict_types=1);

namespace Hyperf\Support\Redis;

use Hyperf\Di\Annotation\Inject;
use Hyperf\Support\Common\RedisModel;
use Hyperf\Extra\Contract\HashServiceInterface;

class RefreshToken extends RedisModel
{
    protected $key = 'refresh-token:';

    /**
     * @Inject()
     * @var HashServiceInterface
     */
    private $hash;

    public function __construct(\Redis $redis = null)
    {
        parent::__construct($redis);
    }

    /**
     * Factory Refresh Token
     * @param string $jti Token ID
     * @param string $ack Ack Code
     * @param int $expires Expires
     * @return mixed
     */
    public function factory(string $jti, string $ack, int $expires)
    {
    }
}