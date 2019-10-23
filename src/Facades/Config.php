<?php
declare(strict_types=1);

namespace Hyperf\Support\Facades;

use Hyperf\Contract\ConfigInterface;
use Hyperf\Support\Common\Facade;

/**
 * Class Config
 * @package Hyperf\Support\Facades
 * @method static mixed get(string $key, mixed $default = null)
 * @method static bool has(string $keys)
 * @method static void set(string $key, mixed $value)
 */
final class Config extends Facade
{
    protected static function getFacadeAccessor()
    {
        return ConfigInterface::class;
    }
}