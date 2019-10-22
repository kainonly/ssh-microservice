<?php
declare(strict_types=1);

namespace Hyperf\Support\Facades;

use Hyperf\Support\Common\Facade;
use Hyperf\Extra\Contract\TokenServiceInterface;

/**
 * Class Token
 * @package Hyperf\Support\Facades
 * @method static \Lcobucci\JWT\Token create(string $scene, array $symbol = [])
 * @method static \Lcobucci\JWT\Token get(string $tokenString)
 * @method static bool verify(string $scene, string $tokenString)
 */
final class Token extends Facade
{
    protected static function getFacadeAccessor()
    {
        return TokenServiceInterface::class;
    }
}