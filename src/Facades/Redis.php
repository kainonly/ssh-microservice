<?php
declare(strict_types=1);

namespace Hyperf\Support\Facades;

use Hyperf\Support\Common\Facade;

/**
 * Class Redis
 * @package Hyperf\Support\Facades
 * @method static string|mixed|bool get($key)
 * @method static bool set($key, $value, $timeout = null)
 * @method static bool setex($key, $ttl, $value)
 */
final class Redis extends Facade
{
    protected static function getFacadeAccessor()
    {
        return \Redis::class;
    }
}